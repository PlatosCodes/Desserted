package gapi

import (
	"context"
	"database/sql"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
	"github.com/PlatosCodes/desserted/backend/pb"
	"github.com/PlatosCodes/desserted/backend/val"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) PlayDessert(ctx context.Context, req *pb.PlayDessertRequest) (*pb.PlayDessertResponse, error) {
	log.Println("request was:", req)
	// Authenticate and authorize the user
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	// Fetch and validate the player game
	playerGame, err := server.Store.GetPlayerGame(ctx, req.GetPlayerGameId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error fetching player game: %v", err)
	}
	if playerGame.PlayerID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "you are not authorized to play in this game")
	}

	cardIDs := req.GetCardIds()

	ingredientsList := []string{}

	for _, cardID := range cardIDs {
		card, err := server.Store.GetCardByID(ctx, cardID)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "not_found":
					return nil, status.Errorf(codes.NotFound, "cannot find card with that ID: %s", err)
				}
			}
			return nil, status.Errorf(codes.Internal, "internal server error: %s", err)
		}
		// Make sure card is in player's hand
		in_hand, err := server.Store.IsCardInPlayerHand(ctx, db.IsCardInPlayerHandParams{
			PlayerGameID: playerGame.PlayerGameID,
			CardID:       cardID,
		})

		if !in_hand {
			return nil, status.Errorf(codes.InvalidArgument, "card played is not in player's hand: %v", err)
		}

		// Record card played (**Make this a db transaction soon**)
		err = server.Store.RecordPlayedCard(ctx, db.RecordPlayedCardParams{
			PlayerGameID: playerGame.PlayerGameID,
			CardID:       cardID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error recording card played: %v", err)
		}
		// Remove card from player's hand (**Make this a db transaction soon**)
		err = server.Store.RemoveCardFromPlayerHand(ctx, db.RemoveCardFromPlayerHandParams{
			PlayerGameID: playerGame.PlayerGameID,
			CardID:       cardID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error removing card player's hand: %v", err)
		}
		ingredientsList = append(ingredientsList, card.Name)
	}

	// Validate the dessert
	err = val.ValidateDessert(req.GetDessertName(), ingredientsList)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid ingredients for dessert: %v", err)
	}

	// Record the dessert played
	dessert, err := server.Store.GetDessertByName(ctx, req.GetDessertName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error fetching dessert ID: %v", err)
	}
	err = server.Store.RecordDessertPlayed(ctx, db.RecordDessertPlayedParams{
		PlayerGameID: req.GetPlayerGameId(),
		DessertID:    dessert.DessertID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error recording dessert played: %v", err)
	}

	// Update player's score
	updated_player_game, err := server.Store.UpdatePlayerScore(ctx, db.UpdatePlayerScoreParams{
		PlayerGameID: playerGame.PlayerGameID,
		PlayerScore: sql.NullInt32{
			Int32: playerGame.PlayerScore.Int32 + dessert.Points,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updated player's score: %s", err)
	}

	winning_condition, err := server.Store.IsGameWon(ctx, playerGame.PlayerGameID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error checking if game was won: %v", err)
	}

	// Convert to pb.PlayDessertResponse
	return &pb.PlayDessertResponse{
		DessertPlayedId: dessert.DessertID,
		PlayerGame: &pb.PlayerGame{
			PlayerGame:   updated_player_game.PlayerGameID,
			PlayerId:     updated_player_game.PlayerID,
			GameId:       updated_player_game.GameID,
			PlayerScore:  updated_player_game.PlayerScore.Int32,
			PlayerStatus: updated_player_game.PlayerStatus.String,
		},
		// Game is won
		GameOver: winning_condition.Bool,
	}, nil

}
