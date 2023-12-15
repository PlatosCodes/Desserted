package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCardByID(t *testing.T) {
	cardIDs := []int64{}

	for i := int64(1); i <= 22; i++ {
		cardIDs = append(cardIDs, i)
	}

	for _, cardID := range cardIDs {
		card, err := testQueries.GetCardByID(context.Background(), cardID)
		require.NoError(t, err)
		require.NotEmpty(t, card)

		require.NotEmpty(t, card.Name)
		require.NotEmpty(t, card.Type)
		require.NotZero(t, card.CardID)

		require.Equal(t, cardID, card.CardID)
	}
}

func TestListCards(t *testing.T) {
	// Retrieve all cards
	cards, err := testQueries.ListCards(context.Background())
	require.NoError(t, err)
	require.Len(t, cards, 50)

	expectedIngredientCount := 29
	expectedSpecialCount := 21

	var ingredientCount, specialCount int

	// Separate cards by type and count them
	for _, card := range cards {
		switch card.Type {
		case "ingredient":
			ingredientCount++
		case "special":
			specialCount++
		default:
			t.Errorf("Unexpected card type: %s", card.Type)
		}
	}

	// Validate counts for each type of card excluding desserts
	require.Equal(t, expectedIngredientCount, ingredientCount)
	require.Equal(t, expectedSpecialCount, specialCount)
}

func TestListCardsByType(t *testing.T) {
	// Define the types of cards you have seeded
	// Modify these types based on your seeded data
	cardTypes := []string{"ingredient", "special"}

	for _, cardType := range cardTypes {
		cardCount := 0
		cardsOfType, err := testQueries.ListCardsByType(context.Background(), cardType)
		require.NoError(t, err)

		for _, card := range cardsOfType {
			cardCount++
			require.Equal(t, cardType, card.Type)
			require.NotEmpty(t, card.Name)
		}
		switch cardType {
		case "ingredient":
			require.Equal(t, cardCount, 29)
		case "special":
			require.Equal(t, cardCount, 21)
		default:
			t.Errorf("Unexpected card type: %s", cardType)
		}

	}
}
