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
	AcceptFriendRequest(ctx context.Context, arg AcceptFriendRequestParams) error
	AcceptGameInvitation(ctx context.Context, arg AcceptGameInvitationParams) error
	ActivateUser(ctx context.Context, id int64) error
	AddCardToPlayerHand(ctx context.Context, arg AddCardToPlayerHandParams) (int64, error)
	AddPlayerToGame(ctx context.Context, arg AddPlayerToGameParams) error
	BlockSession(ctx context.Context, id uuid.UUID) error
	// Checks if all actions for a turn are completed for a player
	CheckAllActionsCompleted(ctx context.Context, playerGameID int64) (sql.NullBool, error)
	// Checks if draw card action for a turn has been completed for a player
	CheckCardDrawn(ctx context.Context, playerGameID int64) (bool, error)
	// Checks if play dessert card action for a turn has been completed for a player
	CheckDessertPlayed(ctx context.Context, playerGameID int64) (bool, error)
	// Checks if play special card action for a turn has been completed for a player
	CheckSpecialCardPlayed(ctx context.Context, playerGameID int64) (bool, error)
	CreateFriendship(ctx context.Context, arg CreateFriendshipParams) (Friend, error)
	CreateGame(ctx context.Context, createdBy int64) (Game, error)
	CreateGameInvitationWithUsername(ctx context.Context, arg CreateGameInvitationWithUsernameParams) error
	CreatePlayerTurnActions(ctx context.Context, playerGameID int64) error
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// Declare the winner of the game
	DeclareWinner(ctx context.Context, gameID int64) (DeclareWinnerRow, error)
	DeleteActivationToken(ctx context.Context, userID int64) error
	DeleteFriendship(ctx context.Context, friendshipID int64) error
	DeleteGameInvitation(ctx context.Context, arg DeleteGameInvitationParams) error
	DeleteUser(ctx context.Context, id int64) error
	DoesInvitationExist(ctx context.Context, arg DoesInvitationExistParams) (bool, error)
	DrawTopCard(ctx context.Context, gameID int64) (int64, error)
	EndGame(ctx context.Context, gameID int64) error
	GetActivationToken(ctx context.Context, activationToken string) (ActivationToken, error)
	// Get card by ID
	GetCardByID(ctx context.Context, cardID int64) (Card, error)
	// Get card by Name
	GetCardByName(ctx context.Context, name string) (Card, error)
	GetDessertByName(ctx context.Context, name string) (Dessert, error)
	GetDessertIDByName(ctx context.Context, name string) (int64, error)
	GetDessertsPlayedByPlayer(ctx context.Context, playerGameID int64) ([]int64, error)
	GetFriendshipByID(ctx context.Context, friendshipID int64) (Friend, error)
	GetGameByID(ctx context.Context, gameID int64) (Game, error)
	GetGameByPlayerGameID(ctx context.Context, playerGameID int64) (GetGameByPlayerGameIDRow, error)
	GetGameDeck(ctx context.Context, gameID int64) (GameDeck, error)
	GetGameScores(ctx context.Context, gameID int64) ([]GetGameScoresRow, error)
	GetPlayedCards(ctx context.Context, playerGameID int64) ([]PlayedCard, error)
	GetPlayerGame(ctx context.Context, playerGameID int64) (PlayerGame, error)
	GetPlayerHand(ctx context.Context, playerGameID int64) ([]GetPlayerHandRow, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	InsertActivationToken(ctx context.Context, arg InsertActivationTokenParams) (ActivationToken, error)
	InsertIntoGameDeck(ctx context.Context, arg InsertIntoGameDeckParams) (int64, error)
	IsCardInPlayerHand(ctx context.Context, arg IsCardInPlayerHandParams) (bool, error)
	IsDeckEmpty(ctx context.Context, gameID int64) (bool, error)
	// Check if a player has reached the winning condition
	IsGameWon(ctx context.Context, playerGameID int64) (sql.NullBool, error)
	IsUserGameCreator(ctx context.Context, arg IsUserGameCreatorParams) (bool, error)
	ListActiveGames(ctx context.Context, arg ListActiveGamesParams) ([]Game, error)
	ListActivePlayerGames(ctx context.Context, playerID int64) ([]ListActivePlayerGamesRow, error)
	// List all cards
	ListCardIDs(ctx context.Context) ([]int64, error)
	// List all cards
	ListCards(ctx context.Context) ([]Card, error)
	// List cards by type
	ListCardsByType(ctx context.Context, type_ string) ([]Card, error)
	ListGameInvitationsForUser(ctx context.Context, inviteePlayerID int64) ([]GameInvitation, error)
	ListGamePlayers(ctx context.Context, arg ListGamePlayersParams) ([]PlayerGame, error)
	ListPendingFriendRequests(ctx context.Context, friendeeID int64) ([]ListPendingFriendRequestsRow, error)
	ListPlayerGames(ctx context.Context, playerID int64) ([]PlayerGame, error)
	ListUserFriends(ctx context.Context, arg ListUserFriendsParams) ([]Friend, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	RecordDessertPlayed(ctx context.Context, arg RecordDessertPlayedParams) error
	RecordPlayedCard(ctx context.Context, arg RecordPlayedCardParams) error
	RemoveCardFromDeck(ctx context.Context, arg RemoveCardFromDeckParams) error
	RemoveCardFromPlayerHand(ctx context.Context, arg RemoveCardFromPlayerHandParams) error
	// Resets the turn actions for a player after their turn
	ResetTurnActions(ctx context.Context, playerGameID int64) error
	StartGame(ctx context.Context, arg StartGameParams) error
	// Updates the card drawn status for a player
	UpdateCardDrawnStatus(ctx context.Context, playerGameID int64) error
	// Updates the dessert played status for a player
	UpdateDessertPlayedStatus(ctx context.Context, playerGameID int64) error
	UpdateGameState(ctx context.Context, arg UpdateGameStateParams) error
	UpdateGameStatus(ctx context.Context, arg UpdateGameStatusParams) error
	UpdatePlayerNumber(ctx context.Context, arg UpdatePlayerNumberParams) error
	UpdatePlayerScore(ctx context.Context, arg UpdatePlayerScoreParams) (PlayerGame, error)
	UpdatePlayerStatus(ctx context.Context, arg UpdatePlayerStatusParams) error
	// Updates the special card played status for a player
	UpdateSpecialCardPlayedStatus(ctx context.Context, playerGameID int64) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
