package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var randSeed struct {
	value int64
	once  sync.Once
}

var randGenerator *rand.Rand

func Rand() *rand.Rand {
	randSeed.once.Do(func() {
		randSeed.value = time.Now().UnixMicro()
		randGenerator = rand.New(rand.NewSource(randSeed.value))
	})
	return randGenerator
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[Rand().Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(10) + strconv.Itoa(time.Now().Nanosecond())

}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomCard() int64 {
	return Rand().Int63n(50) + 1
}

// RandomIngredient generates a random ingredient.
func RandomIngredient() string {
	ingredients := []string{
		"Flour", "Sugar", "Eggs", "Butter", "Cocoa", "Vanilla",
		"Berries", "Cream Cheese", "Saffron", "Honey",
		"Dark Chocolate", "Matcha Powder", "Edible Gold Leaf",
	}
	return ingredients[rand.Intn(len(ingredients))]
}

// RandomDessert generates a random dessert.
func RandomDessertName() string {
	dessertName := []string{
		"Cake", "Pie", "Chocolate Chip Cookies", "Cheesecake", "Tiramisu",
		"Matcha Cake", "Saffron Panna Cotta", "Gourmet Truffles", "Gold Leaf Cupcakes",
	}
	return dessertName[rand.Intn(len(dessertName))]
}

// getRandomPoints returns a random integer between 1 and 10
func RandomPoints() int32 {
	return Rand().Int31n(10) + 1
}
