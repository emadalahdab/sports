package racingboard

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"time"

	"go.uber.org/zap"

	"github.com/robbydyer/sports/internal/board"
	"github.com/robbydyer/sports/internal/logo"
	"github.com/robbydyer/sports/internal/rgbrender"
	scrcnvs "github.com/robbydyer/sports/internal/scrollcanvas"
)

// ScrollRender ...
func (s *RacingBoard) ScrollRender(ctx context.Context, canvas board.Canvas, padding int) (board.Canvas, error) {
	origScrollMode := s.config.ScrollMode.Load()
	origPad := s.config.TightScrollPadding
	defer func() {
		s.config.ScrollMode.Store(origScrollMode)
		s.config.TightScrollPadding = origPad
	}()

	s.config.ScrollMode.Store(true)
	s.config.TightScrollPadding = padding

	return s.render(ctx, canvas)
}

// Render ...
func (s *RacingBoard) Render(ctx context.Context, canvas board.Canvas) error {
	c, err := s.render(ctx, canvas)
	if err != nil {
		return err
	}
	if c != nil {
		defer func() {
			if scr, ok := c.(*scrcnvs.ScrollCanvas); ok {
				s.config.scrollDelay = scr.GetScrollSpeed()
			}
		}()
		return c.Render(ctx)
	}

	return nil
}

// Render ...
func (s *RacingBoard) render(ctx context.Context, canvas board.Canvas) (board.Canvas, error) {
	s.boardCtx, s.boardCancel = context.WithCancel(ctx)

	s.log.Debug("render racing board",
		zap.String("league", s.api.LeagueShortName()),
	)
	if s.leagueLogo == nil {
		var err error
		s.leagueLogo, err = s.api.GetLogo(ctx, canvas.Bounds())
		if err != nil {
			return nil, err
		}
	}

	if len(s.events) < 1 {
		var err error
		s.events, err = s.api.GetScheduledEvents(ctx)
		if err != nil {
			return nil, err
		}
	}

	scheduleWriter, err := s.getScheduleWriter(rgbrender.ZeroedBounds(canvas.Bounds()))
	if err != nil {
		return nil, err
	}

	var scrollCanvas *scrcnvs.ScrollCanvas
	if canvas.Scrollable() && s.config.ScrollMode.Load() {
		base, ok := canvas.(*scrcnvs.ScrollCanvas)
		if !ok {
			return nil, fmt.Errorf("invalid scroll canvas")
		}

		var err error
		scrollCanvas, err = scrcnvs.NewScrollCanvas(base.Matrix, s.log,
			scrcnvs.WithMergePadding(s.config.TightScrollPadding),
		)
		if err != nil {
			return nil, err
		}
		scrollCanvas.SetScrollSpeed(s.config.scrollDelay)
		scrollCanvas.SetScrollDirection(scrcnvs.RightToLeft)
		base.SetScrollSpeed(s.config.scrollDelay)
		go scrollCanvas.MatchScroll(ctx, base)
	}

	s.log.Debug("racing events",
		zap.String("league", s.api.LeagueShortName()),
		zap.Int("number", len(s.events)),
	)

EVENTS:
	for _, event := range s.events {
		select {
		case <-s.boardCtx.Done():
			return nil, context.Canceled
		default:
		}
		img, err := s.renderEvent(s.boardCtx, canvas.Bounds(), event, s.leagueLogo, scheduleWriter)
		if err != nil {
			s.log.Error("failed to render racing event",
				zap.Error(err),
			)
			continue EVENTS
		}

		if scrollCanvas != nil && s.config.ScrollMode.Load() {
			scrollCanvas.AddCanvas(img)
			continue EVENTS
		}

		draw.Draw(canvas, img.Bounds(), img, image.Point{}, draw.Over)

		if err := canvas.Render(s.boardCtx); err != nil {
			s.log.Error("failed to render racing board",
				zap.Error(err),
			)
			continue EVENTS
		}

		if !s.config.ScrollMode.Load() {
			select {
			case <-ctx.Done():
				return nil, context.Canceled
			case <-time.After(s.config.boardDelay):
			}
		}
	}

	if canvas.Scrollable() && scrollCanvas != nil {
		return scrollCanvas, nil
	}

	return nil, nil
}

func (s *RacingBoard) renderEvent(ctx context.Context, bounds image.Rectangle, event *Event, leagueLogo *logo.Logo, scheduleWriter *rgbrender.TextWriter) (draw.Image, error) {
	img := image.NewRGBA(bounds)
	canvasBounds := rgbrender.ZeroedBounds(bounds)

	logoImg, err := leagueLogo.RenderRightAlignedWithEnd(ctx, canvasBounds, (canvasBounds.Max.X-canvasBounds.Min.X)/2)
	if err != nil {
		return nil, err
	}

	pt := image.Pt(logoImg.Bounds().Min.X, logoImg.Bounds().Min.Y)
	draw.Draw(img, logoImg.Bounds(), logoImg, pt, draw.Over)

	gradient := rgbrender.GradientXRectangle(
		canvasBounds,
		0.1,
		color.Black,
		s.log,
	)
	pt = image.Pt(gradient.Bounds().Min.X, gradient.Bounds().Min.Y)
	draw.Draw(img, gradient.Bounds(), gradient, pt, draw.Over)

	event.Date = event.Date.Local()

	tz, _ := event.Date.Zone()
	txt := []string{
		event.Name,
		event.Date.Format("01/02/2006"),
		fmt.Sprintf("%s %s", event.Date.Format("3:04PM"), tz),
	}

	lengths, err := scheduleWriter.MeasureStrings(img, txt)
	if err != nil {
		return nil, err
	}
	max := canvasBounds.Dx() / 2

	for _, length := range lengths {
		if length > max {
			max = length
		}
	}

	s.log.Debug("max racing schedule text length",
		zap.Int("max", max),
		zap.Int("half bounds", canvasBounds.Dy()/2),
	)

	scheduleBounds := image.Rect(
		canvasBounds.Max.X/2,
		canvasBounds.Min.Y,
		(canvasBounds.Max.X/2)+max,
		canvasBounds.Max.Y,
	)

	if err := scheduleWriter.WriteAligned(
		rgbrender.LeftCenter,
		img,
		scheduleBounds,
		txt,
		color.White,
	); err != nil {
		return nil, fmt.Errorf("failed to write schedule: %w", err)
	}

	return img, nil
}
