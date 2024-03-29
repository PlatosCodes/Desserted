package gameservice

type EventType string

const (
	EventTypeCardDrawn         EventType = "CardDrawn"
	EventTypeDessertPlayed     EventType = "DessertPlayed"
	EventTypeSpecialCardPlayed EventType = "SpecialCardPlayed"
	EventTypeScoreUpdate       EventType = "ScoreUpdate"
	EventTypeEndTurn           EventType = "EndTurn"
	EventTypeEndGame           EventType = "EndGame"
)

// Event represents a game event
type Event struct {
	Type EventType
	Data interface{}
}

// CardDrawnData carries data for a card drawn event
type CardDrawnData struct {
	CardID       int64  `json:"card_id"`
	CardName     string `json:"name"`
	PlayerGameID int64  `json:"player_game_id"`
	PlayerHandID int64  `json:"player_hand_id"`
	GameID       int64  `json:"game_id"`
}

// DessertPlayedData carries data for a dessert played event
type DessertPlayedData struct {
	GameID       int64  `json:"game_id"`
	PlayerGameID int64  `json:"player_game_id"`
	PlayerNumber int32  `json:"player_number"`
	DessertName  string `json:"dessert_name"`
	DessertScore int32  `json:"dessert_score"`
	Score        int32  `json:"score"`
	GameOver     bool   `json:"game_over"`
	Success      bool   `json:"success"`
}

// SpecialCardPlayedData carries data for a special card played event
type SpecialCardPlayedData struct {
}

// ScoreUpdateData carries data for a score update event
type ScoreUpdateData struct {
	PlayerGameID int64  `json:"player_game_id"`
	PlayerID     int64  `json:"player_id"`
	GameID       int64  `json:"game_id"`
	PlayerNumber int32  `json:"player_number"`
	PlayerScore  int32  `json:"player_score"`
	PlayerStatus string `json:"player_status"`
}

// EndTurnData carries data for an end turn event
type EndTurnData struct {
	GameID              int64  `json:"game_id"`
	Status              string `json:"status"`
	CreatedBy           int64  `json:"created_by"`
	NumberOfPlayers     int32  `json:"number_of_players"`
	CurrentTurn         int32  `json:"current_turn"`
	CurrentPlayerNumber int32  `json:"current_player_number"`
}

// EndGameData carries data for an end game event
type EndGameData struct {
	GameID              int64  `json:"game_id"`
	Status              string `json:"status"`
	WinningPlayerGameID int64  `json:"winner_player_game_id"`
	WinningPlayerNumber int32  `json:"winner_player_number"`
	WinningScore        int32  `json:"winning_score"`
}
