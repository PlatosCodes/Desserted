package util

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
