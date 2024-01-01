package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddPlayerToGame(t *testing.T) {
	game := createRandomGame(t)
	user := createRandomUser(t)

	add_player_params := AddPlayerToGameParams{
		PlayerID: user.ID,
		GameID:   game.GameID,
	}

	err := testQueries.AddPlayerToGame(context.Background(), add_player_params)
	require.NoError(t, err)

	player_games, err := testQueries.ListPlayerGames(context.Background(), user.ID)
	require.NoError(t, err)

	require.Equal(t, game.GameID, player_games[0].GameID)
	require.Equal(t, user.ID, player_games[0].PlayerID)

}

func TestUpdatePlayerScore(t *testing.T) {
	game := createRandomGame(t)
	user := createRandomUser(t)

	add_player_params := AddPlayerToGameParams{
		PlayerID: user.ID,
		GameID:   game.GameID,
	}

	err := testQueries.AddPlayerToGame(context.Background(), add_player_params)
	require.NoError(t, err)

	player_games, err := testQueries.ListPlayerGames(context.Background(), user.ID)
	require.NoError(t, err)

	player_to_score := player_games[0]

	player_score := player_to_score.PlayerScore

	dessert_points := 10

	updated_score := player_score + int32(dessert_points)

	updated_player, err := testQueries.UpdatePlayerScore(context.Background(), UpdatePlayerScoreParams{
		PlayerScore:  updated_score,
		PlayerGameID: player_to_score.PlayerGameID,
	})
	require.NoError(t, err)

	require.Equal(t, updated_player.PlayerScore, updated_score)
}
