package calendarboard

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"time"

	"go.uber.org/zap"

	"github.com/robbydyer/sports/internal/board"
	"github.com/robbydyer/sports/internal/rgbrender"
	scrcnvs "github.com/robbydyer/sports/internal/scrollcanvas"
)

// ScrollRender ...
func (s *CalendarBoard) ScrollRender(ctx context.Context, canvas board.Canvas, padding int) (board.Canvas, error) {
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
func (s *CalendarBoard) Render(ctx context.Context, canvas board.Canvas) error {
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
func (s *CalendarBoard) render(ctx context.Context, canvas board.Canvas) (board.Canvas, error) {
	s.boardCtx, s.boardCancel = context.WithCancel(ctx)

	events, err := s.api.DailyEvents(ctx, time.Now())
	if err != nil {
		return nil, err
	}

	s.log.Debug("calendar events",
		zap.Int("number", len(events)),
	)

	if len(events) < 1 {
		return nil, nil
	}

	scheduleWriter, err := s.getScheduleWriter(rgbrender.ZeroedBounds(canvas.Bounds()))
	if err != nil {
		return nil, err
	}

	if s.logo == nil {
		var err error
		s.logo, err = s.api.CalendarIcon(ctx, canvas.Bounds())
		if err != nil {
			return nil, err
		}
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

EVENTS:
	for _, event := range events {
		select {
		case <-s.boardCtx.Done():
			return nil, context.Canceled
		default:
		}
		img, err := s.renderEvent(s.boardCtx, canvas.Bounds(), event, scheduleWriter)
		if err != nil {
			s.log.Error("failed to render calendar event",
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
			s.log.Error("failed to render calendar board",
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

func (s *CalendarBoard) renderEvent(ctx context.Context, bounds image.Rectangle, event *Event, writer *rgbrender.TextWriter) (draw.Image, error) {
	img := image.NewRGBA(bounds)
	canvasBounds := rgbrender.ZeroedBounds(bounds)

	logoHeight := int(writer.FontSize * 2.0)
	logoBounds := image.Rect(canvasBounds.Min.X, canvasBounds.Min.Y, canvasBounds.Min.X+logoHeight, canvasBounds.Min.Y+logoHeight)

	dateBounds := image.Rect(canvasBounds.Min.X+logoHeight+2, canvasBounds.Min.Y, canvasBounds.Max.X, canvasBounds.Min.Y+logoHeight)

	titleBounds := image.Rect(canvasBounds.Min.X, canvasBounds.Min.Y+logoHeight+2, canvasBounds.Max.X, canvasBounds.Max.Y)

	logoImg, err := s.logo.RenderLeftAlignedWithStart(ctx, logoBounds, 0)
	if err != nil {
		return nil, err
	}

	pt := image.Pt(logoImg.Bounds().Min.X, logoImg.Bounds().Min.Y)
	draw.Draw(img, logoImg.Bounds(), logoImg, pt, draw.Over)

	if err := writer.WriteAligned(
		rgbrender.CenterCenter,
		img,
		dateBounds,
		[]string{
			event.Time.Format("Mon Jan 2"),
			event.Time.Format("03:04PM"),
		},
		color.White,
	); err != nil {
		return nil, err
	}

	lines, err := writer.BreakText(img, titleBounds.Max.X-titleBounds.Min.X, event.Title)
	if err != nil {
		return nil, err
	}

	maxLines := int(math.Ceil(float64(titleBounds.Dy()) / writer.FontSize))

	if len(lines) > maxLines {
		lines = lines[0:maxLines]
	}

	s.log.Debug("calendar event",
		zap.Strings("titles", lines),
		zap.Int("max lines", maxLines),
		zap.Int("X min", titleBounds.Min.X),
		zap.Int("Y min", titleBounds.Min.Y),
		zap.Int("X max", titleBounds.Max.X),
		zap.Int("Y max", titleBounds.Max.Y),
	)

	if err := writer.WriteAligned(
		rgbrender.LeftBottom,
		img,
		titleBounds,
		lines,
		color.White,
	); err != nil {
		return nil, err
	}

	return img, nil
}
