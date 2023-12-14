package val

import (
	"desserted-backend/util" // Import your utility package for dessert type validation
	"errors"
	"fmt"
)

// ValidateDessert validates if a dessert is valid based on game rules
func ValidateDessert(dessertType string, ingredientCards []string) error {
	// Check if the dessert type is supported
	if !util.IsSupportedDessertType(dessertType) {
		return errors.New("unsupported dessert type")
	}

	// Define the required ingredients for each dessert type
	requiredIngredients := map[string][]string{
		util.Cake:                 {"Flour", "Sugar", "Eggs"},
		util.Pie:                  {"Flour", "Butter", "Berries"},
		util.ChocolateChipCookies: {"Flour", "Sugar", "Chocolate"},
		util.Cheesecake:           {"Cream Cheese", "Eggs", "Vanilla"},
		util.Tiramisu:             {"Coffee", "Cream Cheese", "Cocoa"},
		util.MatchaCake:           {"Flour", "Matcha Powder", "Eggs"},
		util.SaffronPannaCotta:    {"Cream", "Saffron", "Sugar"},
		util.GourmetTruffles:      {"Dark Chocolate", "Cream", "Honey"},
		util.GoldLeafCupcakes:     {"Flour", "Sugar", "Edible Gold Leaf"},
	}

	// Check if the ingredient cards match the required ingredients
	required, found := requiredIngredients[dessertType]
	if !found {
		return errors.New("unknown dessert type")
	}

	// Convert ingredientCards to a map for easier validation
	ingredientMap := make(map[string]bool)
	for _, ingredient := range ingredientCards {
		ingredientMap[ingredient] = true
	}

	// Check if all required ingredients are present
	for _, ingredient := range required {
		if !ingredientMap[ingredient] {
			return fmt.Errorf("missing required ingredient: %s", ingredient)
		}
	}

	// Check for any extra ingredients
	for ingredient := range ingredientMap {
		if !util.IsSupportedIngredientType(ingredient) {
			return fmt.Errorf("unsupported ingredient: %s", ingredient)
		}
	}

	return nil
}
