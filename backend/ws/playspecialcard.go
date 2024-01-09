package ws

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
)

type SpecialCardPayload struct {
	PlayerGameID int64  `json:"player_game_id"`
	CardType     string `json:"card_type"` // Type of special card
}

func (c *Client) handlePlaySpecialCard(payload json.RawMessage) {
	var specialCardPayload SpecialCardPayload

	if err := json.Unmarshal(payload, &specialCardPayload); err != nil {
		log.Printf("Error unmarshaling special card payload: %v", err)
		return
	}
	log.Println("SPECIAL CARD PAYLOAD:", specialCardPayload)

	c.mutex.Lock()
	defer c.mutex.Unlock()

	ctx := context.Background()

	switch specialCardPayload.CardType {
	case "RefreshPantry":
		err := c.store.RefreshPlayerPantryTx(ctx, specialCardPayload.PlayerGameID)
		if err != nil {
			log.Printf("Error refreshing player hand: %v", err)
		} else {
			// Fetch the new hand and send it to the client
			newHand, err := c.store.GetPlayerHand(ctx, specialCardPayload.PlayerGameID)
			if err != nil {
				log.Printf("Error fetching new hand: %v", err)
			} else {
				// Send new hand to the client
				log.Println(newHand)
				c.sendUpdatedHand(newHand)
			}
		}

	// case "StealCard", "InstantBake", "Sabotage":
	// Implement other special cards here...

	default:
		log.Printf("Unknown special card type: %s", specialCardPayload.CardType)
	}

	// Notify client about the action taken
	// ...
}

func (c *Client) sendUpdatedHand(newHand []db.GetPlayerHandRow) {
	handUpdate := struct {
		Type string                `json:"type"`
		Hand []db.GetPlayerHandRow `json:"hand"`
	}{
		Type: "refreshPantry",
		Hand: newHand,
	}

	msg, err := json.Marshal(handUpdate)
	if err != nil {
		log.Printf("Error marshaling hand update: %v", err)
		return
	}

	c.send <- msg
}
