package game

import (
	"errors"
	"fmt"

	"github.com/PlatosCodes/desserted/backend/util"
)

// DessertPointsMap holds the points for each dessert type.
var DessertPointsMap = map[string]int32{
	"Cake":                           10,
	"Pie":                            15,
	"Chocolate Chip Cookies":         20,
	"Cheesecake":                     25,
	"Marble Cake":                    30,
	"Triple Chocolate Brownies":      35,
	"Gourmet Truffles":               40,
	"Raspberry Chocolate Cheesecake": 45,
	"Gold Leaf Cupcakes":             50,
}

// GetDessertPoints returns the points for a given dessert.
func GetDessertPoints(dessertName string) (int32, bool) {
	points, exists := DessertPointsMap[dessertName]
	return points, exists
}

// CalculateDessertScore calculates the score for a dessert based on ingredients and special cards.
func CalculateDessertScore(dessertName string, ingredientsList []string, specialCard string) (int32, error) {
	var doublePointsMultiplier int32 = 1
	var extraPoints int32
	baseScore, ok := GetDessertPoints(dessertName)
	if !ok {
		return 0, fmt.Errorf("unknown dessert type: %s", dessertName)
	}

	if specialCard != "" {
		doublePointsMultiplier, extraPoints = ProcessSpecialCards(specialCard)
	}

	return (baseScore * doublePointsMultiplier) + extraPoints, nil
}

// ProcessSpecialCards processes the special cards and returns score modifiers.
func ProcessSpecialCards(specialCard string) (int32, int32) {
	var doublePointsMultiplier int32 = 1
	var extraPoints int32 = 0

	switch specialCard {
	case "Double Points":
		doublePointsMultiplier = 2
	case "Glass of Milk":
		extraPoints += 3
	case "Mystery Ingredient":
		extraPoints += util.RandomPoints()
	}

	return doublePointsMultiplier, extraPoints
}

// ValidateDessert validates if a dessert is valid based on game rules
func ValidateDessert(dessertName string, ingredientCards []string) error {
	// Check if the dessert type is supported
	if !IsSupportedDessertType(dessertName) {
		return errors.New("unsupported dessert type")
	}

	requiredIngredients, err := GetRequiredIngredientsForDessert(dessertName)
	if err != nil {
		return err // Unsupported or invalid dessert
	}

	// Convert ingredientCards to a map for easier validation
	ingredientMap := make(map[string]bool)
	for _, ingredient := range ingredientCards {
		ingredientMap[ingredient] = true
	}

	// Check if all required ingredients are present
	for _, ingredient := range requiredIngredients {
		if !ingredientMap[ingredient] {
			return fmt.Errorf("missing required ingredient: %s", ingredient)
		}
	}

	// Check for any extra ingredients
	for ingredient := range ingredientMap {
		if !IsSupportedIngredientType(ingredient) {
			return fmt.Errorf("%s is an unsupported ingredient for dessert type: %s", dessertName, ingredient)
		}
	}

	return nil
}
