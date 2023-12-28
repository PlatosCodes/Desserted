// rpc_get_game_by_id.go
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

// GetGameByID retrieves the details of a game by its ID.
func (server *Server) GetGameByID(ctx context.Context, req *pb.GetGameByIDRequest) (*pb.GetGameByIDResponse, error) {
	// Authenticate and authorize the user
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// ******* CONTINUE FROM HERE : MAKE SURE PLAYER IS PLAYER IN GAME -> RETURN GAME //

	gameID := req.GetGameId()
	if gameID < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "not a valid user ID: %d", gameID)
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

	for _, player := range players {
		if authPayload.UserID == player.PlayerID {
			game_player = true
		}
	}

	if !game_player {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get game you are not a player in")
	}

	// Fetch the game data
	game, err := server.Store.GetGameByID(ctx, req.GetGameId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving game data: %v", err)
	}

	// Convert game data to protobuf message
	pbGame := &pb.Game{
		GameId:        game.GameID,
		Status:        game.Status,
		CreatedBy:     game.CreatedBy,
		CurrentTurn:   int32(game.CurrentTurn),
		CurrentPlayer: game.CurrentPlayerID.Int64,
		StartTime:     timestamppb.New(game.StartTime),
		EndTime:       nil,
	}

	// Construct response
	response := &pb.GetGameByIDResponse{
		Game: pbGame,
	}

	log.Println("Response for get game by id:", response)

	return response, nil
}
