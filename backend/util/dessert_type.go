package util

// Constants for all supported dessert types
const (
	Cake                 = "Cake"
	Pie                  = "Pie"
	ChocolateChipCookies = "Chocolate Chip Cookies"
	Cheesecake           = "Cheesecake"
	Tiramisu             = "Tiramisu"
	MatchaCake           = "Matcha Cake"
	SaffronPannaCotta    = "Saffron Panna Cotta"
	GourmetTruffles      = "Gourmet Truffles"
	GoldLeafCupcakes     = "Gold Leaf Cupcakes"
)

// IsSupportedDessertType returns true if the dessert type is supported
func IsSupportedDessertType(dessertType string) bool {
	switch dessertType {
	case Cake, Pie, ChocolateChipCookies, Cheesecake, Tiramisu, MatchaCake, SaffronPannaCotta, GourmetTruffles, GoldLeafCupcakes:
		return true
	}
	return false
}
