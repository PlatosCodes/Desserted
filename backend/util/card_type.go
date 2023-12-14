package util

// Constants for all supported card types
const (
	Ingredient = "ingredient"
	Special    = "special"
)

// IsSupportedCardType returns true if the card type is supported
func IsSupportedCardype(card_type string) bool {
	switch card_type {
	case Ingredient, Special:
		return true
	}
	return false
}
