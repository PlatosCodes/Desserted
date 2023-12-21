package gapi

import (
	"context"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DrawCard handles the request to draw a card from the deck in a game session.
func (server *Server) DrawCard(ctx context.Context, req *pb.DrawCardRequest) (*pb.DrawCardResponse, error) {
	// Authenticate the user and ensure they are authorized to perform this action.
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	gameID, playerGameID := req.GetGameId(), req.GetPlayerGameId()

	// Validate the player's participation in the game.
	playerGame, err := server.Store.GetPlayerGame(ctx, playerGameID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error validating player in game: %v", err)
	}
	if playerGame.PlayerID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "you are not authorized to draw cards in this game")
	}

	// Perform the card drawing within a database transaction to ensure data consistency.
	drawn_card, err := server.Store.DrawCard(ctx, db.DrawCardTxParams{
		GameID:   gameID,
		PlayerID: playerGame.PlayerID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error drawing top card: %v", err)
	}

	// Construct and return the response to the client.
	return &pb.DrawCardResponse{
		CardId: drawn_card,
	}, nil
}
