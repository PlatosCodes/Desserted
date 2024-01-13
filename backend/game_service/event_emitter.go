package gameservice

import (
	"log"
	"sync"
)

var (
	eventEmitter = make(chan Event, 100)
	once         sync.Once
)

func EmitEvent(event Event) {
	log.Printf("Emitting event: %v", event)

	eventEmitter <- event
}

func StartEventDispatcher(handleEvent func(Event)) {
	once.Do(func() {
		go func() {
			for event := range eventEmitter {
				handleEvent(event)
			}
		}()
	})
}

func (s *GameService) emitDessertPlayEvent(gameID, playerGameID int64, playerNumber int32, dessertName string, score int32, success, gameOver bool) {
	event := Event{
		Type: EventTypeDessertPlayed,
		Data: DessertPlayedData{
			GameID:       gameID,
			PlayerGameID: playerGameID,
			PlayerNumber: playerNumber,
			DessertName:  dessertName,
			Score:        score,
			DessertScore: score,
			GameOver:     gameOver,
			Success:      success,
		},
	}
	EmitEvent(event)
}
