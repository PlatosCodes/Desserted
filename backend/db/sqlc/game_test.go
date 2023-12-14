package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"desserted-backend/util"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomGame(t *testing.T) Game {
	user := createRandomUser(t)
	userID := user.ID

	game, err := testQueries.CreateGame(context.Background(), userID)
	require.NoError(t, err)
	require.NotEmpty(t, game)

	require.Equal(t, userID, game.CreatedBy)

	require.NotZero(t, game.ID)
	require.NotZero(t, game.CreatedAt)

	return game
}

func TestCreateGame(t *testing.T) {
	createRandomGame(t)
}

func createRandomDessert(t *testing.T) Dessert {
	ingredient := util.RandomSupportedIngredient()   // Replace with your random ingredient generation logic
	dessertType := util.RandomSupportedDessertType() // Replace with your random dessert type generation logic

	dessertParams := CreateDessertParams{
		Name:        util.RandomDessertName(),
		Points:      util.RandomDessertPoints(),
		Ingredient:  ingredient,
		DessertType: dessertType,
	}

	dessert, err := testQueries.CreateDessert(context.Background(), dessertParams)
	require.NoError(t, err)
	require.NotEmpty(t, dessert)

	// Add any additional assertions here

	return dessert
}

func TestCreateDessert(t *testing.T) {
	createRandomDessert(t)
}

func TestGetGameByID(t *testing.T) {
	randomGame := createRandomGame(t)

	game, err := testQueries.GetGameByID(context.Background(), randomGame.ID)

	require.NoError(t, err)
	require.NotEmpty(t, game)

	require.Equal(t, randomGame.ID, game.ID)
	require.Equal(t, randomGame.Status, game.Status)
	require.Equal(t, randomGame.CreatedBy, game.CreatedBy)

	require.NotZero(t, game.ID)
	require.NotZero(t, game.CreatedAt)

	require.WithinDuration(t, randomGame.CreatedAt, game.CreatedAt, time.Second)

}

func TestUpdateGameStatus(t *testing.T) {
	randomGame := createRandomGame(t)

	err := testQueries.EndGame(context.Background(), randomGame.ID)
	require.NoError(t, err)

	game, err := testQueries.GetGameByID(context.Background(), randomGame.ID)
	require.NoError(t, err)
	require.NotEmpty(t, game)

	require.Equal(t, game.Status.String, "complete")

	require.NotZero(t, game.EndedAt)
	require.WithinDuration(t, time.Now(), game.EndedAt.Time, time.Second)

}

func TestDrawCard(t *testing.T) {
	card, err := testQueries.DrawCard(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, card)
}

func TestCreateDessert(t *testing.T) {
	randomGame := createRandomGame(t)
	new_user := createRandomUser(t)

	nullUserType := sql.NullInt64{
		Int64: new_user.ID,
		Valid: true,
	}

	nullGameIDType := sql.NullInt64{
		Int64: randomGame.ID,
		Valid: true,
	}

	// Add the test player to the players table
	err := testQueries.AddPlayerToGame(context.Background(), AddPlayerToGameParams{
		UserID: nullUserType,
		GameID: nullGameIDType,
	})

	require.NoError(t, err)

	// Define the required "ingredient" cards for a dessert
	ingredientCards := []string{"Flour", "Sugar", "Eggs"}

	// Draw the required "ingredient" cards
	for _, cardName := range ingredientCards {
		card, err := testQueries.DrawCard(context.Background())
		require.NoError(t, err)
		require.NotNil(t, card)
		require.Equal(t, cardName, card.Name)
	}

	// Play the "ingredient" cards to create the dessert
	for _, cardName := range ingredientCards {
		err := testQueries.PlayDessert(context.Background(), PlayDessertParams{
			ArrayAppend: cardName,
			UserID:      sql.NullInt32{Int32: userID, Valid: true},
			GameID:      sql.NullInt32{Int32: gameID, Valid: true},
		})
		require.NoError(t, err)
	}

	// Fetch the player's played cards to check if the dessert card was added
	player, err := testQueries.GetPlayerByID(context.Background(), userID, gameID)
	require.NoError(t, err)
	require.NotNil(t, player)

	// Define the dessert card's name based on the combination of "ingredient" cards
	dessertName := "Cake"

	// Assert that the dessert card is in the played_cards array
	assert.Contains(t, player.PlayedCards, dessertName)
}
