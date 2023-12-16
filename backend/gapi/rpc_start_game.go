package gapi

import (
	"context"
	"database/sql"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// StartGame starts a game session.
func (server *Server) StartGame(ctx context.Context, req *pb.StartGameRequest) (*pb.StartGameResponse, error) {
	// Authenticate and authorize the user
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// Check if the user is the game creator
	game, err := server.Store.GetGameByID(ctx, req.GetGameId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "game not found")
		}
		return nil, status.Errorf(codes.Internal, "error fetching game: %v", err)
	}

	if game.CreatedBy != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "only the game creator can start the game")
	}

	players, err := server.Store.ListGamePlayers(ctx, db.ListGamePlayersParams{
		GameID: req.GetGameId(),
		Limit:  4,
		Offset: 0,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list game players: %v", err)
	}

	var player_ids []int64
	for _, player := range players {
		player_ids = append(player_ids, player.PlayerID)
	}

	// Start the game using the start game transactional method
	startGameResult, err := server.Store.StartGameTx(ctx, db.StartGameTxParams{
		GameID:    req.GetGameId(),
		PlayerIDs: player_ids,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to start game: %v", err)
	}

	var endTime *timestamppb.Timestamp

	if game.EndTime.Valid {
		endTime = timestamppb.New(game.EndTime.Time)
	} else {
		endTime = nil
	}

	return &pb.StartGameResponse{
		Game: &pb.Game{
			GameId:    startGameResult.Game.GameID,
			Status:    startGameResult.Game.Status,
			CreatedBy: startGameResult.Game.CreatedBy,
			StartTime: timestamppb.New(startGameResult.Game.StartTime),
			EndTime:   endTime,
		},
	}, nil
}
