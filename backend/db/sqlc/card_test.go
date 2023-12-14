package db

import (
	"context"
	"database/sql"
	"testing"

	"desserted-backend/util"

	"github.com/stretchr/testify/require"
)

func TestGetCardByID(t *testing.T) {
	cardID := util.RandomCard()

	card, err := testQueries.GetCardByID(context.Background(), cardID)
	require.NoError(t, err)
	require.NotEmpty(t, card)

	require.NotZero(t, card.ID)
	require.Equal(t, cardID, card.ID)
}

// func TestListCards
func TestListCards(t *testing.T) {
	cards, err := testQueries.ListCards(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, cards)
}

func TestListCardsByType(t *testing.T) {
	cardTypes := []string{util.Ingredient, util.Special}

	for _, cardType := range cardTypes {
		// Create a sql.NullString with the cardType value
		nullCardType := sql.NullString{
			String: cardType,
			Valid:  true,
		}

		cards, err := testQueries.ListCardsByType(context.Background(), nullCardType)
		require.NoError(t, err)
		require.NotEmpty(t, cards)

		// Check card types for each card in the result
		for _, card := range cards {
			require.Equal(t, cardType, card.Type.String)
		}
	}
}
