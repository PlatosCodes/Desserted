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
