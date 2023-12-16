package gapi

import (
	"context"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Will use this functionality in the future if players want to "find an open game"
// Requires adding a private/public column to games
func (server *Server) AddPlayerToGame(ctx context.Context, req *pb.AddPlayerToGameRequest) (*emptypb.Empty, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	player_to_add := req.GetPlayerId()
	if authPayload.UserID != player_to_add {
		return nil, status.Errorf(codes.PermissionDenied, "you are not authorized to add a game for this user")
	}

	arg := db.AddPlayerToGameParams{
		PlayerID: player_to_add,
		GameID:   req.GetGameId(),
	}

	err = server.Store.AddPlayerToGame(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "player has already joined game: %s", err)
			default:
				return nil, status.Errorf(codes.Internal, "database error in adding player to game: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	return &emptypb.Empty{}, nil
}
