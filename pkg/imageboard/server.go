package imageboard

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/twitchtv/twirp"

	pb "github.com/robbydyer/sports/internal/proto/imageboard"
)

// Server ...
type Server struct {
	board *ImageBoard
}

// GetRPCHandler ...
func (i *ImageBoard) GetRPCHandler() (string, http.Handler) {
	return i.rpcServer.PathPrefix(), i.rpcServer
}

// SetStatus ...
func (s *Server) SetStatus(ctx context.Context, req *pb.SetStatusReq) (*emptypb.Empty, error) {
	if req.Status == nil {
		return &emptypb.Empty{}, twirp.NewError(twirp.InvalidArgument, "nil status sent")
	}

	s.board.config.Enabled.Store(req.Status.Enabled)

	return &emptypb.Empty{}, nil
}

// GetStatus ...
func (s *Server) GetStatus(ctx context.Context, req *emptypb.Empty) (*pb.StatusResp, error) {
	return &pb.StatusResp{
		Status: &pb.Status{
			Enabled: s.board.config.Enabled.Load(),
		},
	}, nil
}

// Jump ...
func (s *Server) Jump(ctx context.Context, req *pb.JumpReq) (*emptypb.Empty, error) {
	i := s.board
	i.jumpLock.Lock()
	defer i.jumpLock.Unlock()

	// Clear the channel
	select {
	case <-i.jumpTo:
	default:
	}

	select {
	case i.jumpTo <- req.Name:
	case <-time.After(5 * time.Second):
		return &emptypb.Empty{}, twirp.InternalError("timed out attempting image jump")
	}

	if i.jumpTo != nil {
		if err := i.jumper(i.Name()); err != nil {
			i.log.Error("failed to jump to image board",
				zap.Error(err),
				zap.String("file name", req.Name),
			)
			return &emptypb.Empty{}, twirp.InternalError("failed to jump to image board")
		}
	}

	return &emptypb.Empty{}, nil
}