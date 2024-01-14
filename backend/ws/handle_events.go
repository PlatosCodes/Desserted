package ws

import (
	"encoding/json"
	"log"

	gameservice "github.com/PlatosCodes/desserted/backend/game_service"
)

func (h *Hub) HandleGameEvents() {
	gameservice.StartEventDispatcher(func(event gameservice.Event) {
		switch event.Type {
		case gameservice.EventTypeCardDrawn:
			h.handleCardDrawnEvent(event.Data.(gameservice.CardDrawnData))
		case gameservice.EventTypeDessertPlayed:
			h.handleDessertPlayedEvent(event.Data.(gameservice.DessertPlayedData))
		// case gameservice.EventTypeSpecialCardPlayed:
		// 	h.handleSpecialCardPlayedEvent(event.Data.(gameservice.SpecialCardPlayedData))
		case gameservice.EventTypeScoreUpdate:
			h.handleScoreUpdateEvent(event.Data.([]gameservice.ScoreUpdateData))
		case gameservice.EventTypeEndTurn:
			h.handleEndTurnEvent(event.Data.(gameservice.EndTurnData))
		case gameservice.EventTypeEndGame:
			h.handleEndGameEvent(event.Data.(gameservice.EndGameData))

		}
	})
}

func (h *Hub) handleCardDrawnEvent(data gameservice.CardDrawnData) {

	type CardData struct {
		CardID       int64  `json:"card_id"`
		CardName     string `json:"name"`
		PlayerGameID int64  `json:"player_game_id"`
		PlayerHandID int64  `json:"player_hand_id"`
	}

	msg := struct {
		Type string   `json:"type"`
		Data CardData `json:"cardDrawnData"`
	}{
		Type: "drawCardResponse",
		Data: CardData{
			CardID:       data.CardID,
			CardName:     data.CardName,
			PlayerGameID: data.PlayerGameID,
			PlayerHandID: data.PlayerHandID,
		},
	}

	serializedMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling dessert played data: %v", err)
		return
	}

	// Broadcast the message to all clients in the game
	h.sendToClient(data.PlayerGameID, data.GameID, serializedMsg)
}

func (h *Hub) handleDessertPlayedEvent(data gameservice.DessertPlayedData) {
	msg := struct {
		Type string                        `json:"type"`
		Data gameservice.DessertPlayedData `json:"data"`
	}{
		Type: "dessertPlayedUpdate",
		Data: data,
	}

	serializedMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling dessert played data: %v", err)
		return
	}

	// Broadcast the message to all clients in the game
	h.broadcastMessage(data.GameID, serializedMsg)
}

// func (h *Hub) handleSpecialCardPlayedEvent(data gameservice.SpecialCardPlayedData) {
// 	// Convert to WebSocket message and send to relevant clients
// 	_, err := json.Marshal(data) // Customize as needed
// 	if err != nil {
// 		log.Printf("Error marshaling dessert played data: %v", err)
// 		return
// 	}
// 	// h.broadcastMessage(msg)
// }

func (h *Hub) handleScoreUpdateEvent(scoreData []gameservice.ScoreUpdateData) {
	msg := struct {
		Type string                        `json:"type"`
		Data []gameservice.ScoreUpdateData `json:"score_data"`
	}{
		Type: "scoreUpdate",
		Data: scoreData,
	}

	serializedMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling score update data: %v", err)
		return
	}

	h.broadcastMessage(scoreData[0].GameID, serializedMsg)
}

func (h *Hub) handleEndTurnEvent(endTurnData gameservice.EndTurnData) {
	msg := struct {
		Type string                  `json:"type"`
		Data gameservice.EndTurnData `json:"end_turn_data"`
	}{
		Type: "endTurnUpdate",
		Data: endTurnData,
	}

	serializedMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling end turn data: %v", err)
		return
	}

	h.broadcastMessage(endTurnData.GameID, serializedMsg)
}

func (h *Hub) handleEndGameEvent(endGameData gameservice.EndGameData) {
	msg := struct {
		Type string                  `json:"type"`
		Data gameservice.EndGameData `json:"end_game_data"`
	}{
		Type: "endGameUpdate",
		Data: endGameData,
	}

	serializedMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling end game data: %v", err)
		return
	}

	h.broadcastMessage(endGameData.GameID, serializedMsg)
}
