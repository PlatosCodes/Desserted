// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	AcceptGameInvitation(ctx context.Context, arg AcceptGameInvitationParams) error
	AddCardToPlayerHand(ctx context.Context, arg AddCardToPlayerHandParams) error
	AddPlayerToGame(ctx context.Context, arg AddPlayerToGameParams) error
	CreateFriendship(ctx context.Context, arg CreateFriendshipParams) (Friend, error)
	CreateGame(ctx context.Context, createdBy int64) (Game, error)
	CreateGameInvitation(ctx context.Context, arg CreateGameInvitationParams) error
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// Declare the winner of the game
	DeclareWinner(ctx context.Context, gameID int64) (int64, error)
	DeleteFriendship(ctx context.Context, friendshipID int64) error
	DeleteGameInvitation(ctx context.Context, arg DeleteGameInvitationParams) error
	DeleteUser(ctx context.Context, id int64) error
	DoesInvitationExist(ctx context.Context, arg DoesInvitationExistParams) (bool, error)
	DrawTopCard(ctx context.Context, gameID int64) (int64, error)
	EndGame(ctx context.Context, gameID int64) error
	// Get card by ID
	GetCardByID(ctx context.Context, cardID int64) (Card, error)
	GetDessertByName(ctx context.Context, name string) (Dessert, error)
	GetDessertIDByName(ctx context.Context, name string) (int64, error)
	GetDessertsPlayedByPlayer(ctx context.Context, playerGameID int64) ([]int64, error)
	GetFriendshipByID(ctx context.Context, friendshipID int64) (Friend, error)
	GetGameByID(ctx context.Context, gameID int64) (Game, error)
	GetGameDeck(ctx context.Context, gameID int64) (GameDeck, error)
	GetPlayedCards(ctx context.Context, playerGameID int64) ([]PlayedCard, error)
	GetPlayerGame(ctx context.Context, playerGameID int64) (PlayerGame, error)
	GetPlayerHand(ctx context.Context, playerGameID int64) ([]GetPlayerHandRow, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	InsertIntoGameDeck(ctx context.Context, arg InsertIntoGameDeckParams) (int64, error)
	IsCardInPlayerHand(ctx context.Context, arg IsCardInPlayerHandParams) (bool, error)
	IsDeckEmpty(ctx context.Context, gameID int64) (bool, error)
	// Check if a player has reached the winning condition
	IsGameWon(ctx context.Context, playerGameID int64) (sql.NullBool, error)
	IsUserGameCreator(ctx context.Context, arg IsUserGameCreatorParams) (bool, error)
	ListActiveGames(ctx context.Context, arg ListActiveGamesParams) ([]Game, error)
	// List all cards
	ListCardIDs(ctx context.Context) ([]int64, error)
	// List all cards
	ListCards(ctx context.Context) ([]Card, error)
	// List cards by type
	ListCardsByType(ctx context.Context, type_ string) ([]Card, error)
	ListGameInvitationsForUser(ctx context.Context, inviteeUsername string) ([]GameInvitation, error)
	ListGamePlayers(ctx context.Context, arg ListGamePlayersParams) ([]PlayerGame, error)
	ListPlayerGames(ctx context.Context, playerID int64) ([]PlayerGame, error)
	ListUserFriends(ctx context.Context, arg ListUserFriendsParams) ([]Friend, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	RecordDessertPlayed(ctx context.Context, arg RecordDessertPlayedParams) error
	RecordPlayedCard(ctx context.Context, arg RecordPlayedCardParams) error
	RemoveCardFromDeck(ctx context.Context, arg RemoveCardFromDeckParams) error
	RemoveCardFromPlayerHand(ctx context.Context, arg RemoveCardFromPlayerHandParams) error
	StartGame(ctx context.Context, gameID int64) error
	UpdateGameStatus(ctx context.Context, arg UpdateGameStatusParams) error
	UpdatePlayerScore(ctx context.Context, arg UpdatePlayerScoreParams) (PlayerGame, error)
	UpdatePlayerStatus(ctx context.Context, arg UpdatePlayerStatusParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
