package model

import (
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
			Level: levelPlay,
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
		Level: levelPlay,
	}
}
