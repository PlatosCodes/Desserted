package util

// Constants for all supported ingredient types
const (
	Flour          = "Flour"
	Sugar          = "Sugar"
	Eggs           = "Eggs"
	Butter         = "Butter"
	Chocolate      = "Chocolate"
	Vanilla        = "Vanilla"
	Berries        = "Berries"
	CreamCheese    = "Cream Cheese"
	Honey          = "Honey"
	DarkChocolate  = "Dark Chocolate"
	EdibleGoldLeaf = "Edible Gold Leaf"
)

// IsSupportedIngredientType returns true if the ingredient type is supported
func IsSupportedIngredientType(ingredientType string) bool {
	switch ingredientType {
	case Flour, Sugar, Eggs, Butter, Vanilla, Berries, CreamCheese, Honey, Chocolate, EdibleGoldLeaf:
		return true
	}
	return false
}
