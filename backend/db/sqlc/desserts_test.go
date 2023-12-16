package db

import (
	"context"
	"testing"

	"github.com/PlatosCodes/desserted/backend/util"
	"github.com/PlatosCodes/desserted/backend/val"

	"github.com/stretchr/testify/require"
)

func TestRecordDessertPlayed(t *testing.T) {
	player1GameID, _, _ := createActiveGame(t)
	dessertName := util.RandomDessertName()

	dessert_id, err := testQueries.GetDessertIDByName(context.Background(), dessertName)
	require.NoError(t, err)

	err = testQueries.RecordDessertPlayed(context.Background(), RecordDessertPlayedParams{
		PlayerGameID: player1GameID,
		DessertID:    dessert_id,
	})
	require.NoError(t, err)
}

func TestGetDessertsPlayedByPlayer(t *testing.T) {
	playerID, dessertID := createPlayedDessert(t)

	desserts, err := testQueries.GetDessertsPlayedByPlayer(context.Background(), playerID)
	require.NoError(t, err)

	require.Len(t, desserts, 1)
	require.Equal(t, desserts[0], dessertID)

}

func createPlayedDessert(t *testing.T) (playedDessertPlayerID int64, playedDessertID int64) {
	player1GameID, _, _ := createActiveGame(t)
	// Cake is made up of Flour: (id 1-6), Sugar: (id 7-11), Eggs: (id 12-15)

	ingredientCardIDs := []int64{1, 7, 12}
	dessertName := "Cake"

	// // Add flour, sugar and eggs to hand
	// for _, cardID := range ingredientCardIDs {
	// 	err := testQueries.AddCardToPlayerHand(context.Background(), AddCardToPlayerHandParams{
	// 		PlayerGameID: player1GameID,
	// 		CardID:       cardID,
	// 	})
	// 	require.NoError(t, err)
	// }
	var ingredients []string
	for _, cardID := range ingredientCardIDs {
		card, err := testQueries.GetCardByID(context.Background(), cardID)
		require.NoError(t, err)

		ingredients = append(ingredients, card.Name)
	}

	err := val.ValidateDessert(dessertName, ingredients)
	require.NoError(t, err)

	dessert_id, err := testQueries.GetDessertIDByName(context.Background(), dessertName)
	require.NoError(t, err)

	err = testQueries.RecordDessertPlayed(context.Background(), RecordDessertPlayedParams{
		PlayerGameID: player1GameID,
		DessertID:    dessert_id,
	})
	require.NoError(t, err)

	return player1GameID, dessert_id
}

func TestCreatePlayedDessert(t *testing.T) {
	createPlayedDessert(t)
}
