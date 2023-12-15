package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomGame(t *testing.T) Game {
	user := createRandomUser(t)

	game, err := testQueries.CreateGame(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, game)

	require.Equal(t, game.CreatedBy, user.ID)

	require.NotZero(t, game.GameID)
	require.NotZero(t, game.CreatedBy)
	require.True(t, game.EndTime.Time.IsZero())

	return game
}

func TestCreateGame(t *testing.T) {
	createRandomGame(t)
}

func TestGetGameByID(t *testing.T) {
	game := createRandomGame(t)

	gotGame, err := testQueries.GetGameByID(context.Background(), game.GameID)
	require.NoError(t, err)

	require.Equal(t, game.GameID, gotGame.GameID)
	require.Equal(t, game.EndTime, gotGame.EndTime)
	require.Equal(t, game.CreatedBy, gotGame.CreatedBy)
	require.Equal(t, game.StartTime, gotGame.StartTime)
	require.Equal(t, game.Status, gotGame.Status)
}

func TestEndGame(t *testing.T) {
	game := createRandomGame(t)

	err := testQueries.EndGame(context.Background(), game.GameID)
	require.NoError(t, err)

	gotGame, err := testQueries.GetGameByID(context.Background(), game.GameID)
	require.NoError(t, err)

	require.Equal(t, game.GameID, gotGame.GameID)
	require.Equal(t, game.CreatedBy, gotGame.CreatedBy)
	require.Equal(t, game.StartTime, gotGame.StartTime)
	require.Equal(t, "completed", gotGame.Status)

	require.NotEqual(t, game.Status, gotGame.Status)
	require.NotEqual(t, game.EndTime, gotGame.EndTime)

	require.WithinDuration(t, time.Now(), gotGame.EndTime.Time, time.Second)
}

func TestListActiveGames(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomGame(t)
	}

	active_game_params := ListActiveGamesParams{
		Limit:  5,
		Offset: 0,
	}

	games, err := testQueries.ListActiveGames(context.Background(), active_game_params)
	require.NoError(t, err)
	require.Len(t, games, 5)

	for _, game := range games {
		require.NotEmpty(t, game)
	}
}

func TestUpdateGameStatus(t *testing.T) {
	game := createRandomGame(t)

	update_game_status_params := UpdateGameStatusParams{
		Status: "suspended",
		GameID: game.GameID,
	}

	err := testQueries.UpdateGameStatus(context.Background(), update_game_status_params)
	require.NoError(t, err)

	gotGame, err := testQueries.GetGameByID(context.Background(), game.GameID)
	require.NoError(t, err)
	require.Equal(t, "suspended", gotGame.Status)

}

func TestDeclareWinner(t *testing.T) {

}

func TestListGamePlayers(t *testing.T) {
	game := createRandomGame(t)

	add_player_params := AddPlayerToGameParams{
		PlayerID: game.CreatedBy,
		GameID:   game.GameID,
	}

	err := testQueries.AddPlayerToGame(context.Background(), add_player_params)
	require.NoError(t, err)

	user1 := createRandomUser(t)
	add_player1_params := AddPlayerToGameParams{
		PlayerID: user1.ID,
		GameID:   game.GameID,
	}

	err = testQueries.AddPlayerToGame(context.Background(), add_player1_params)
	require.NoError(t, err)

	user2 := createRandomUser(t)
	add_player2_params := AddPlayerToGameParams{
		PlayerID: user2.ID,
		GameID:   game.GameID,
	}

	err = testQueries.AddPlayerToGame(context.Background(), add_player2_params)
	require.NoError(t, err)

	game_params := ListGamePlayersParams{
		GameID: game.GameID,
		Limit:  10,
		Offset: 0,
	}

	player_games, err := testQueries.ListGamePlayers(context.Background(), game_params)
	require.NoError(t, err)
	require.Len(t, player_games, 3)
}
