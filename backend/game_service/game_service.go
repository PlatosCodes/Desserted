package gameservice

import (
	db "github.com/PlatosCodes/desserted/backend/db/sqlc"
)

type GameService struct {
	store db.Store // Interface to interact with db operations
}

// NewGameService creates a new instance of GameService with the provided store
func NewGameService(store db.Store) *GameService {
	return &GameService{
		store: store,
	}
}
