package gapi

import (
	"context"

	"github.com/PlatosCodes/desserted/backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetPlayerHand retrieves the cards in a player's hand for a given game.
func (server *Server) GetPlayerHand(ctx context.Context, req *pb.GetPlayerHandRequest) (*pb.GetPlayerHandResponse, error) {
	// Authenticate and authorize the user
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// Verify that the player is part of the game
	playerGame, err := server.Store.GetPlayerGame(ctx, req.GetPlayerGameId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// Ensure that the requestor is the player in question
	if playerGame.PlayerID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "you can only view your own hand")
	}

	// Fetch the player's hand
	playerHand, err := server.Store.GetPlayerHand(ctx, req.GetPlayerGameId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving player hand: %v", err)
	}

	// Convert to protobuf response
	var pbPlayerHand []*pb.PlayerHand
	for _, card := range playerHand {
		pbPlayerHand = append(pbPlayerHand, &pb.PlayerHand{
			PlayerHandId: card.PlayerHandID,
			PlayerGameId: card.PlayerGameID,
			CardId:       card.CardID,
			CardName:     card.Name,
		})
	}

	return &pb.GetPlayerHandResponse{
		PlayerHand: pbPlayerHand,
	}, nil
}
