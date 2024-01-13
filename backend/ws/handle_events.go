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
		case gameservice.EventTypeSpecialCardPlayed:
			h.handleSpecialCardPlayedEvent(event.Data.(gameservice.SpecialCardPlayedData))
		case gameservice.EventTypeScoreUpdate:
			h.handleScoreUpdateEvent(event.Data.([]gameservice.ScoreUpdateData))
		case gameservice.EventTypeEndTurn:
			h.handleEndTurnEvent(event.Data.(gameservice.EndTurnData))
		}
	})
}

func (h *Hub) handleCardDrawnEvent(data gameservice.CardDrawnData) {
	// Convert to WebSocket message and send to relevant clients
	msg, err := json.Marshal(data) // Customize as needed
	if err != nil {
		log.Printf("Error marshaling card drawn data: %v", err)
		return
	}
	h.messageQueue.Enqueue(msg) // Enqueue the message
}

func (h *Hub) handleDessertPlayedEvent(data gameservice.DessertPlayedData) {
	// Create a message structure as per your frontend requirements
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

func (h *Hub) handleSpecialCardPlayedEvent(data gameservice.SpecialCardPlayedData) {
	// Convert to WebSocket message and send to relevant clients
	_, err := json.Marshal(data) // Customize as needed
	if err != nil {
		log.Printf("Error marshaling dessert played data: %v", err)
		return
	}
	// h.broadcastMessage(msg)
}

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
		log.Printf("Error marshaling dessert played data: %v", err)
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
		log.Printf("Error marshaling dessert played data: %v", err)
		return
	}

	h.broadcastMessage(endTurnData.GameID, serializedMsg)
}
