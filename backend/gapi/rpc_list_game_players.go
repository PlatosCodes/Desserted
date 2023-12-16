package gapi

import (
	"context"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListGamePlayers(ctx context.Context, req *pb.ListGamePlayersRequest) (*pb.ListGamePlayersResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	gameID := req.GetGameId()
	if gameID < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "not a valid user ID: %d", req.GetGameId())
	}

	players, err := server.Store.ListGamePlayers(ctx, db.ListGamePlayersParams{
		GameID: gameID,
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "case_not_found":
				return nil, status.Errorf(codes.NotFound, "game with that ID does not exist: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
	}

	game_player := false

	player_games := []*pb.PlayerGame{}

	for _, player := range players {
		if authPayload.UserID == player.PlayerID {
			game_player = true
		}

		player_games = append(player_games, &pb.PlayerGame{
			PlayerGame:   player.PlayerGameID,
			PlayerId:     player.PlayerID,
			GameId:       player.GameID,
			PlayerScore:  player.PlayerScore.Int64,
			PlayerStatus: player.PlayerStatus.String,
			HandCards:    player.HandCards.String,
			PlayedCards:  player.PlayedCards.String,
		})

	}

	if !game_player {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get game you are not a player in")
	}

	rsp := &pb.ListGamePlayersResponse{
		Players: player_games,
	}

	return rsp, nil
}
