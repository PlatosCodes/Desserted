package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
)

type SpecialCardPayload struct {
	PlayerGameID int64  `json:"player_game_id"`
	CardType     string `json:"card_type"` // Type of special card
	CardID       int64  `json:"card_id"`
}

func (c *Client) handlePlaySpecialCard(payload json.RawMessage) {
	var specialCardPayload SpecialCardPayload

	if err := json.Unmarshal(payload, &specialCardPayload); err != nil {
		log.Printf("Error unmarshaling special card payload: %v", err)
		c.sendErrorMessage(err.Error())
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	ctx := context.Background()

	card, err := c.store.GetCardByID(ctx, specialCardPayload.CardID)
	if err != nil || card.Name != specialCardPayload.CardType {
		errr := errors.New("the card submitted does not match the special card type selected")
		log.Printf("Error in validating special card: %v", errr)
		log.Printf("card sent was: %s and card in db is: %s", specialCardPayload.CardType, card.Name)
		// Optionally, send an error message back to the client
		c.sendErrorMessage("Failed to validate special card: " + errr.Error())
		return
	}

	inHand, err := c.store.IsCardInPlayerHand(ctx, db.IsCardInPlayerHandParams{
		PlayerGameID: specialCardPayload.PlayerGameID,
		CardID:       specialCardPayload.CardID,
	})
	if err != nil || !inHand {
		log.Printf("Error in validating special card in player's hand: %v", err)
		// Optionally, send an error message back to the client
		c.sendErrorMessage("error in validating special card in player's hand: " + err.Error())
		return
	}

	switch specialCardPayload.CardType {
	case "Refresh Pantry":
		err := c.store.RefreshPlayerPantryTx(ctx, specialCardPayload.PlayerGameID, specialCardPayload.CardID)
		if err != nil {
			log.Printf("Error refreshing player hand: %v", err)
		} else {
			// Fetch the new hand and send it to the client
			newHand, err := c.store.GetPlayerHand(ctx, specialCardPayload.PlayerGameID)
			if err != nil {
				log.Printf("Error fetching new hand: %v", err)
				c.sendErrorMessage(err.Error())
			} else {
				// Send new hand to the client
				log.Println(newHand)
				c.sendUpdatedHand(newHand)
			}
		}
	case "Steal Card":
		targetPlayerID, stolenCard, err := c.stealRandomCard(ctx, specialCardPayload.PlayerGameID, specialCardPayload.CardID)
		if err != nil {
			log.Printf("Error in stealing card: %v", err)
			// Optionally, send an error message back to the client
			c.sendErrorMessage("Failed to steal card: " + err.Error())
			return
		}
		c.notifyPlayersAboutSteal(specialCardPayload.PlayerGameID, targetPlayerID, c.gameID, stolenCard)

	// case "StealCard", "InstantBake", "Sabotage":
	// Implement other special cards here...

	default:
		log.Printf("Unknown special card type: %s", specialCardPayload.CardType)
	}

	// Notify client about the action taken
	// ...

	// Check if all actions are completed
	completed, err := c.store.CheckAllActionsCompleted(ctx, specialCardPayload.PlayerGameID)
	if err != nil {
		log.Printf("Error checking actions completed: %v", err)

	}

	log.Println("COMPLETED STATUS IS:", completed)

	if completed.Bool {
		endTurnPayload := EndTurnPayload{
			GameID:       c.gameID,
			PlayerGameID: specialCardPayload.PlayerGameID,
		}

		marshaledPayload, err := json.Marshal(endTurnPayload)
		if err != nil {
			log.Printf("Error marshaling endTurnPayload response: %v", err)
			return
		}

		c.handleEndTurn(marshaledPayload)
	}
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

func (c *Client) stealRandomCard(ctx context.Context, playerGameID int64, cardID int64) (int64, db.Card, error) {
	rsp, err := c.store.StealRandomCardFromPlayerTx(ctx, playerGameID, cardID)
	if err != nil {
		return 0, db.Card{}, err
	}
	stolenCard, err := c.store.GetCardByID(ctx, rsp.StolenCardID)
	if err != nil {
		return 0, db.Card{}, err
	}
	return rsp.TargetPlayerID, stolenCard, nil
}

func (c *Client) notifyPlayersAboutSteal(playerGameID int64, targetPlayerID int64, gameID int64, card db.Card) {
	// Construct the detailed WebSocket message for the involved players
	detailedNotification := struct {
		Type           string  `json:"type"`
		PlayerGameID   int64   `json:"playerGameID"`
		TargetPlayerID int64   `json:"targetPlayerID"`
		StolenCard     db.Card `json:"card"`
	}{
		Type:           "stealCardDetailedNotification",
		PlayerGameID:   playerGameID,
		TargetPlayerID: targetPlayerID,
		StolenCard:     card,
	}

	detailedNotificationJSON, err := json.Marshal(detailedNotification)
	if err != nil {
		log.Println("Error marshaling detailed notification for steal card: %w", err)
	}

	// Send detailed notifications
	c.hub.sendToClient(playerGameID, gameID, detailedNotificationJSON)
	c.hub.sendToClient(targetPlayerID, gameID, detailedNotificationJSON)

	// Notify other players with a generic message
	genericNotification := struct {
		Type             string `json:"type"`
		InitiatingPlayer int64  `json:"initiatingPlayer"`
		AffectedPlayer   int64  `json:"affectedPlayer"`
		NotificationText string `json:"notificationText"`
	}{
		Type:             "stealCardGenericNotification",
		InitiatingPlayer: playerGameID,
		AffectedPlayer:   targetPlayerID,
		NotificationText: fmt.Sprintf("Player %d stole a card from Player %d", playerGameID, targetPlayerID),
	}

	// Send generic notifications
	c.hub.broadcastExcept([]int64{playerGameID, targetPlayerID}, gameID, genericNotification)
}
