package util

// Constants for all supported ingredient types
const (
	Flour          = "Flour"
	Sugar          = "Sugar"
	Eggs           = "Eggs"
	Butter         = "Butter"
	Milk           = "Milk"
	Chocolate      = "Chocolate"
	Vanilla        = "Vanilla"
	Berries        = "Berries"
	Nuts           = "Nuts"
	CreamCheese    = "Cream Cheese"
	Saffron        = "Saffron"
	Honey          = "Honey"
	DarkChocolate  = "Dark Chocolate"
	MatchaPowder   = "Matcha Powder"
	EdibleGoldLeaf = "Edible Gold Leaf"
)

// IsSupportedIngredientType returns true if the ingredient type is supported
func IsSupportedIngredientType(ingredientType string) bool {
	switch ingredientType {
	case Flour, Sugar, Eggs, Butter, Milk, Chocolate, Vanilla, Berries, Nuts, CreamCheese, Saffron, Honey, DarkChocolate, MatchaPowder, EdibleGoldLeaf:
		return true
	}
	return false
}
