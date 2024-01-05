package model

import (
	"errors"

	pb "github.com/sangharsh/color-sort/gen/modelpb"
	"google.golang.org/protobuf/proto"
)

func NewLevelPlay(level *pb.LevelState) *pb.LevelPlay {
	return &pb.LevelPlay{
		Moves:        []*pb.PourSuccessResponse{},
		CurrentState: proto.Clone(level).(*pb.LevelState),
	}
}

func Pour(pourReq *pb.PourRequest, levelPlay *pb.LevelPlay) *pb.PourResponse {
	success, err := pour(levelPlay.GetCurrentState(), int(pourReq.GetSrc()), int(pourReq.GetDst()))
	if !success || err != nil {
		return &pb.PourResponse{
			Status: &pb.PourResponse_Err{
				Err: err.Error(),
			},
			Level: levelPlay.GetCurrentState(),
		}
	}

	pourRes := &pb.PourSuccessResponse{
		Src:            pourReq.GetSrc(),
		Dst:            pourReq.GetDst(),
		NumItemsPoured: 1, // Functionality is limited to 1 right now
	}

	levelPlay.Moves = append(levelPlay.Moves, pourRes)
	return &pb.PourResponse{
		Status: &pb.PourResponse_Response{
			Response: pourRes,
		},
		Level: levelPlay.GetCurrentState(),
	}
}

func Undo(undoReq *pb.UndoRequest, levelPlay *pb.LevelPlay) (*pb.LevelState, error) {
	moves := levelPlay.GetMoves()
	if len(moves) == 0 {
		return nil, errors.New("no moves to undo")
	}
	lastMove := moves[len(moves)-1]
	success, err := undo(levelPlay.GetCurrentState(), lastMove)
	if err != nil || !success {
		return nil, err
	}
	return levelPlay.GetCurrentState(), nil
}
