package db

import (
	"context"
	"testing"

	"github.com/PlatosCodes/desserted/backend/util"

	"github.com/stretchr/testify/require"
)

func createActiveGame(t *testing.T) (player1_game_id int64, player2_game_id int64, game Game) {
	game = createRandomGame(t)
	player2 := createRandomUser(t)

	add_player_params := AddPlayerToGameParams{
		PlayerID: game.CreatedBy,
		GameID:   game.GameID,
	}

	err := testQueries.AddPlayerToGame(context.Background(), add_player_params)
	require.NoError(t, err)

	add_player2_params := AddPlayerToGameParams{
		PlayerID: player2.ID,
		GameID:   game.GameID,
	}

	err = testQueries.AddPlayerToGame(context.Background(), add_player2_params)
	require.NoError(t, err)

	players, err := testQueries.ListGamePlayers(context.Background(), ListGamePlayersParams{
		GameID: game.GameID,
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)

	return players[0].PlayerGameID, players[1].PlayerGameID, game
}

func createRandomPlayerHand(t *testing.T, player1_game_id int64) []GetPlayerHandRow {
	for i := 0; i < 5; i++ {
		cardID := util.RandomCard()

		addCardParams := AddCardToPlayerHandParams{
			PlayerGameID: player1_game_id,
			CardID:       cardID,
		}

		err := testQueries.AddCardToPlayerHand(context.Background(), addCardParams)
		require.NoError(t, err)
	}

	player1Hand, err := testQueries.GetPlayerHand(context.Background(), player1_game_id)
	require.NoError(t, err)
	require.NotEmpty(t, player1Hand)
	require.Len(t, player1Hand, 5)

	for _, hand := range player1Hand {
		require.NotEmpty(t, hand.Name, "Card name should not be empty")
	}

	return player1Hand
}

func TestCreateActiveGame(t *testing.T) {
	createActiveGame(t)
}

func TestCreateRandomPlayerHand(t *testing.T) {
	player1_game_id, _, _ := createActiveGame(t)
	createRandomPlayerHand(t, player1_game_id)
}

func TestAddCardToHand(t *testing.T) {
	player1_game_id, _, _ := createActiveGame(t)

	cardID := util.RandomCard()

	add_card_params := AddCardToPlayerHandParams{
		PlayerGameID: player1_game_id,
		CardID:       cardID,
	}

	err := testQueries.AddCardToPlayerHand(context.Background(), add_card_params)
	require.NoError(t, err)

	player1_hand, err := testQueries.GetPlayerHand(context.Background(), player1_game_id)
	require.NoError(t, err)
	require.NotEmpty(t, player1_hand)
	require.Len(t, player1_hand, 1)

	require.Equal(t, cardID, player1_hand[0].CardID)
}

func TestGetPlayerHand(t *testing.T) {
	player1_game_id, _, _ := createActiveGame(t)

	cardID := util.RandomCard()

	add_card_params := AddCardToPlayerHandParams{
		PlayerGameID: player1_game_id,
		CardID:       cardID,
	}

	err := testQueries.AddCardToPlayerHand(context.Background(), add_card_params)
	require.NoError(t, err)

	player1_hand, err := testQueries.GetPlayerHand(context.Background(), player1_game_id)
	require.NoError(t, err)
	require.NotEmpty(t, player1_hand)
	require.Len(t, player1_hand, 1)

	for i, card_id := range player1_hand {
		require.Equal(t, card_id, player1_hand[i])
	}

}

func TestRecordPlayerCard(t *testing.T) {
	player1_game_id, _, _ := createActiveGame(t)
	player1_hand := createRandomPlayerHand(t, player1_game_id)

	played_card := RecordPlayedCardParams{
		PlayerGameID: player1_game_id,
		CardID:       player1_hand[0].CardID,
	}

	err := testQueries.RecordPlayedCard(context.Background(), played_card)
	require.NoError(t, err)

	played_cards, err := testQueries.GetPlayedCards(context.Background(), player1_game_id)
	require.NoError(t, err)

	require.Len(t, played_cards, 1)

	require.Equal(t, played_card.CardID, played_cards[0].CardID)
}

func TestRemovePlayerCardHand(t *testing.T) {
	player1_game_id, _, _ := createActiveGame(t)
	player_hand := createRandomPlayerHand(t, player1_game_id)

	play_card := player_hand[0].CardID

	err := testQueries.RemoveCardFromPlayerHand(context.Background(), RemoveCardFromPlayerHandParams{
		PlayerGameID: player1_game_id,
		CardID:       play_card,
	})
	require.NoError(t, err)

	player_updated_hand, err := testQueries.GetPlayerHand(context.Background(), player1_game_id)
	require.NoError(t, err)

	require.Len(t, player_updated_hand, 4)

	for i := 0; i < len(player_updated_hand); i++ {
		require.NotEqual(t, play_card, player_updated_hand[i])
	}
}
