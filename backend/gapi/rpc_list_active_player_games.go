// gapi/list_active_player_games.go

package gapi

import (
	"context"
	"log"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListActivePlayerGames lists active games for a player.
func (server *Server) ListActivePlayerGames(ctx context.Context, req *pb.ListPlayerGamesRequest) (*pb.ListPlayerGamesResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// Fetch active player games from the database
	activeGames, err := server.Store.ListActivePlayerGames(ctx, authPayload.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch active games: %v", err)
	}

	log.Println("Active Games:", activeGames)

	// Prepare the response
	playerGames := make([]*pb.PlayerGame, len(activeGames))
	for i, game := range activeGames {
		playerGames[i] = &pb.PlayerGame{
			PlayerGameId:    game.PlayerGameID,
			PlayerId:        game.PlayerID,
			GameId:          game.GameID,
			NumberOfPlayers: game.NumberOfPlayers,
			PlayerNumber:    game.PlayerNumber.Int32,
			PlayerScore:     game.PlayerScore,
			PlayerStatus:    game.PlayerStatus,
			CreatedBy:       game.CreatedBy,
			Status:          game.Status,
		}
	}

	log.Println(&pb.ListPlayerGamesResponse{PlayerGames: playerGames})

	return &pb.ListPlayerGamesResponse{PlayerGames: playerGames}, nil
}
