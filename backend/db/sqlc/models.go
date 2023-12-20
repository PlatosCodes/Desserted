// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Details of each card used in the game
type Card struct {
	CardID int64  `json:"card_id"`
	Type   string `json:"type"`
	Name   string `json:"name"`
}

type Dessert struct {
	DessertID int64  `json:"dessert_id"`
	Name      string `json:"name"`
	Points    int32  `json:"points"`
}

type DessertPlayed struct {
	DessertPlayedID int64          `json:"dessert_played_id"`
	PlayerGameID    int64          `json:"player_game_id"`
	DessertID       int64          `json:"dessert_id"`
	IconPath        sql.NullString `json:"icon_path"`
	Timestamp       time.Time      `json:"timestamp"`
}

type Friend struct {
	FriendshipID int64     `json:"friendship_id"`
	FrienderID   int64     `json:"friender_id"`
	FriendeeID   int64     `json:"friendee_id"`
	FriendedAt   time.Time `json:"friended_at"`
}

// Represents a game session
type Game struct {
	GameID    int64        `json:"game_id"`
	Status    string       `json:"status"`
	CreatedBy int64        `json:"created_by"`
	StartTime time.Time    `json:"start_time"`
	EndTime   sql.NullTime `json:"end_time"`
}

type GameDeck struct {
	GameDeckID int64 `json:"game_deck_id"`
	GameID     int64 `json:"game_id"`
	CardID     int64 `json:"card_id"`
	OrderIndex int32 `json:"order_index"`
}

type GameInvitation struct {
	GameInvitationID int64     `json:"game_invitation_id"`
	InviterPlayerID  int64     `json:"inviter_player_id"`
	InviteeUsername  string    `json:"invitee_username"`
	GameID           int64     `json:"game_id"`
	Timestamp        time.Time `json:"timestamp"`
}

type PlayedCard struct {
	PlayedCardID int64     `json:"played_card_id"`
	PlayerGameID int64     `json:"player_game_id"`
	CardID       int64     `json:"card_id"`
	PlayTime     time.Time `json:"play_time"`
}

// Associates users with their game sessions and tracks their progress
type PlayerGame struct {
	PlayerGameID int64          `json:"player_game_id"`
	PlayerID     int64          `json:"player_id"`
	GameID       int64          `json:"game_id"`
	PlayerScore  sql.NullInt32  `json:"player_score"`
	PlayerStatus sql.NullString `json:"player_status"`
}

type PlayerHand struct {
	PlayerHandID int64 `json:"player_hand_id"`
	PlayerGameID int64 `json:"player_game_id"`
	CardID       int64 `json:"card_id"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

// Stores user account information
type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          []byte    `json:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	TotalScore        int32     `json:"total_score"`
	TotalWins         int32     `json:"total_wins"`
	TotalLosses       int32     `json:"total_losses"`
}
