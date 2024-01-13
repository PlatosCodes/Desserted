package gameservice

import (
	"context"
	"fmt"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
)

func (s *GameService) UpdateScoresForAllPlayers(ctx context.Context, gameID int64) error {

	playerData, err := s.store.ListGamePlayers(ctx, db.ListGamePlayersParams{
		GameID: gameID,
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		log.Printf("Error getting player data for score update: %v", err)
		return fmt.Errorf("error getting player data for score update: %v", err)
	}

	var scoreUpdateData []ScoreUpdateData
	for _, player := range playerData {
		scoreUpdateData = append(scoreUpdateData, ScoreUpdateData{
			PlayerGameID: player.PlayerGameID,
			PlayerID:     player.PlayerID,
			GameID:       player.GameID,
			PlayerNumber: player.PlayerNumber.Int32,
			PlayerScore:  player.PlayerScore,
			PlayerStatus: player.PlayerStatus,
		})
	}

	scoreUpdateEvent := Event{
		Type: EventTypeScoreUpdate,
		Data: scoreUpdateData,
	}

	EmitEvent(scoreUpdateEvent)

	return nil
}
