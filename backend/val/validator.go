package val

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"

	"github.com/PlatosCodes/desserted/backend/util" // Import your utility package for dessert type validation
)

var isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
var isValidRecipeName = regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`).MatchString

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lower case letters, digits or underscores")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	_, err := mail.ParseAddress(value)
	if err != nil {
		return fmt.Errorf("not a valid email address")
	}
	return nil
}

// ValidateDessert validates if a dessert is valid based on game rules
func ValidateDessert(dessertType string, ingredientCards []string) error {
	// Check if the dessert type is supported
	if !util.IsSupportedDessertType(dessertType) {
		return errors.New("unsupported dessert type")
	}

	// Define the required ingredients for each dessert type
	requiredIngredients := map[string][]string{
		util.Cake:                    {"Flour", "Sugar", "Eggs"},
		util.Pie:                     {"Flour", "Butter", "Berries"},
		util.ChocolateChipCookies:    {"Flour", "Sugar", "Chocolate"},
		util.Cheesecake:              {"Cream Cheese", "Eggs", "Vanilla"},
		util.MarbleCake:              {"Flour", "Sugar", "Eggs", "Butter", "Vanilla", "Chocolate"},
		util.TripleChocolateBrownies: {"Chocolate", "Butter", "Sugar", "Flour", "Eggs"},
		util.GourmetTruffles:         {"Chocolate", "Cream Cheese", "Honey"},
		util.RaspberryChocCheesecake: {"Cream Cheese", "Eggs", "Sugar", "Vanilla", "Chocolate", "Berries"},
		util.GoldLeafCupcakes:        {"Flour", "Sugar", "Butter", "Edible Gold Leaf"},
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
