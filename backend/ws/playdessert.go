package ws

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
)

// PlayDessertPayload mirrors the PlayDessertRequest from gRPC
type PlayDessertPayload struct {
	PlayerGameID int64   `json:"player_game_id"`
	DessertName  string  `json:"dessert_name"`
	CardIDs      []int64 `json:"card_ids"`
}

// Custom struct that mirrors pb.PlayDessertResponse but is safe for marshaling
type PlayDessertResponseJSON struct {
	DessertPlayedId int64           `json:"dessert_played_id"`
	PlayerGame      *PlayerGameJSON `json:"player_game"`
	GameOver        bool            `json:"game_over"`
}

type PlayerGameJSON struct {
	PlayerGame   int64  `json:"player_game_id"`
	PlayerId     int64  `json:"player_id"`
	GameId       int64  `json:"game_id"`
	PlayerScore  int32  `json:"player_score"`
	PlayerStatus string `json:"player_status"`
}

func (c *Client) handlePlayDessert(payload json.RawMessage) {
	var playDessertPayload PlayDessertPayload
	if err := json.Unmarshal(payload, &playDessertPayload); err != nil {
		log.Printf("Error unmarshaling play dessert payload: %v", err)
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	ctx := context.Background()
	result, err := c.store.PlayDessertTx(ctx, db.PlayDessertTxParams{
		PlayerGameID: playDessertPayload.PlayerGameID,
		DessertName:  playDessertPayload.DessertName,
		CardIDs:      playDessertPayload.CardIDs,
	})

	if err != nil {
		log.Println("Error processing PlayDessertTx:", err)
		return
	}

	// Fetch updated scores
	updatedPlayers, err := c.store.ListGamePlayers(ctx, db.ListGamePlayersParams{
		GameID: result.PlayerGame.GameID,
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		log.Printf("Error fetching updated game scores: %v", err)
		return
	}

	log.Println("SCORE UPDATE: ", updatedPlayers, "FOR PLAYER: ", playDessertPayload.PlayerGameID)
	// Prepare and send score update message to all clients
	scoreUpdateMsg := prepareScoreUpdateMessage(updatedPlayers)
	c.hub.broadcast <- scoreUpdateMsg
}

func prepareScoreUpdateMessage(players []db.PlayerGame) []byte {
	// Define a struct for the message
	type ScoreUpdate struct {
		Type    string           `json:"type"`
		Players []PlayerGameJSON `json:"players"`
	}

	var scoreData []PlayerGameJSON
	for _, player := range players {
		scoreData = append(scoreData, PlayerGameJSON{
			PlayerGame:   player.PlayerGameID,
			PlayerId:     player.PlayerID,
			GameId:       player.GameID,
			PlayerScore:  player.PlayerScore,
			PlayerStatus: player.PlayerStatus,
		})
	}

	updateMsg := ScoreUpdate{
		Type:    "scoreUpdate",
		Players: scoreData,
	}

	msg, err := json.Marshal(updateMsg)
	if err != nil {
		log.Printf("Error marshaling score update message: %v", err)
		return nil
	}

	return msg
}
