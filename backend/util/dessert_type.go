package util

import (
	"errors"
	"fmt"
)

// Constants for all supported dessert types
const (
	Cake                    = "Cake"
	Pie                     = "Pie"
	ChocolateChipCookies    = "Chocolate Chip Cookies"
	Cheesecake              = "Cheesecake"
	MarbleCake              = "Marble Cake"
	TripleChocolateBrownies = "Triple Chocolate Brownies"
	GourmetTruffles         = "Gourmet Truffles"
	RaspberryChocCheesecake = "Raspberry Chocolate Cheesecake"
	GoldLeafCupcakes        = "Gold Leaf Cupcakes"
)

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
