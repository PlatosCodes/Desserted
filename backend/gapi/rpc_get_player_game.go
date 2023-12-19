// rpc_get_player_game.go
package gapi

import (
	"context"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetPlayerGame retrieves the game data for a given player.
func (server *Server) GetPlayerGame(ctx context.Context, req *pb.GetPlayerGameRequest) (*pb.GetPlayerGameResponse, error) {
	// Authenticate and authorize the user
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// Fetch the player game data
	playerGame, err := server.Store.GetPlayerGame(ctx, req.GetPlayerGameId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving player game data: %v", err)
	}

	// Ensure that the requestor is the player in question or has appropriate permissions
	if playerGame.PlayerID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "you do not have permission to view this game data")
	}

	pbPlayer := &pb.PlayerGame{
		PlayerGame:   playerGame.PlayerGameID,
		PlayerId:     playerGame.PlayerID,
		GameId:       playerGame.GameID,
		PlayerScore:  playerGame.PlayerScore.Int32,
		PlayerStatus: playerGame.PlayerStatus.String,
	}

	// Construct response
	response := &pb.GetPlayerGameResponse{
		Player: pbPlayer,
	}

	return response, nil
}
