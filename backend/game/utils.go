package game

import (
	"errors"
	"fmt"
)

// IsSupportedCardType returns true if the card type is supported
func IsSupportedCardype(card_type string) bool {
	switch card_type {
	case Ingredient, Special:
		return true
	}
	return false
}

// IsSupportedIngredientType returns true if the ingredient type is supported
func IsSupportedIngredientType(ingredientType string) bool {
	switch ingredientType {
	case Flour, Sugar, Eggs, Butter, Vanilla, Berries, CreamCheese, Honey, Chocolate, EdibleGoldLeaf:
		return true
	}
	return false
}

// IsSupportedSpecialCardType returns true if the special card type is supported
func IsSupportedSpecialCardType(specialCardType string) bool {
	switch specialCardType {
	case WildcardIngredient, DoublePoints, MysteryIngredient, StealCard, RefreshPantry, GlassOfMilk:
		return true
	}
	return false
}

// IsSupportedDessertType returns true if the dessert type is supported
func IsSupportedDessertType(dessertType string) bool {
	switch dessertType {
	case Cake, Pie, ChocolateChipCookies, Cheesecake, GourmetTruffles, GoldLeafCupcakes, TripleChocolateBrownies, RaspberryChocCheesecake, MarbleCake: // Include all supported dessert types
		return true
	}
	return false
}

// GetRequiredIngredientsForDessert returns the required ingredients for a given dessert
func GetRequiredIngredientsForDessert(dessertType string) ([]string, error) {
	if !IsSupportedDessertType(dessertType) {
		return nil, errors.New("unsupported dessert type")
	}

	requiredIngredients := map[string][]string{
		Cake:                    {"Flour", "Sugar", "Eggs"},
		Pie:                     {"Flour", "Butter", "Berries"},
		ChocolateChipCookies:    {"Flour", "Sugar", "Chocolate"},
		Cheesecake:              {"Cream Cheese", "Eggs", "Vanilla"},
		MarbleCake:              {"Flour", "Sugar", "Eggs", "Butter", "Vanilla", "Chocolate"},
		TripleChocolateBrownies: {"Chocolate", "Butter", "Sugar", "Flour", "Eggs"},
		GourmetTruffles:         {"Chocolate", "Cream Cheese", "Honey"},
		RaspberryChocCheesecake: {"Cream Cheese", "Eggs", "Sugar", "Vanilla", "Chocolate", "Raspberries"},
		GoldLeafCupcakes:        {"Flour", "Sugar", "Butter", "Edible Gold Leaf"},
	}

	ingredients, exists := requiredIngredients[dessertType]
	if !exists {
		return nil, fmt.Errorf("unknown dessert type: %s", dessertType)
	}

	return ingredients, nil
}

func IsGameWon(score int32) bool {
	return score > WinningScore
}
