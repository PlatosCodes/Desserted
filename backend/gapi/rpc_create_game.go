package gapi

import (
	"context"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateGame(ctx context.Context, req *pb.CreateGameRequest) (*pb.CreateGameResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	log.Println("req:", req.GetCreatedBy())

	creatorID := req.GetCreatedBy()
	if creatorID < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "not a valid user ID: %d", req.GetCreatedBy())
	}

	if authPayload.UserID != req.GetCreatedBy() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create game for another user")
	}

	game, err := server.Store.CreateGame(ctx, creatorID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	err = server.Store.AddPlayerToGame(ctx, db.AddPlayerToGameParams{
		PlayerID: creatorID,
		GameID:   game.GameID,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "player has already joined game: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	var endTime *timestamppb.Timestamp

	if game.EndTime.Valid {
		endTime = timestamppb.New(game.EndTime.Time)
	} else {
		endTime = nil
	}

	rsp := &pb.CreateGameResponse{
		Game: &pb.Game{
			GameId:    game.GameID,
			CreatedBy: game.CreatedBy,
			Status:    game.Status,
			StartTime: timestamppb.New(game.StartTime),
			EndTime:   endTime,
		},
	}

	return rsp, nil
}
