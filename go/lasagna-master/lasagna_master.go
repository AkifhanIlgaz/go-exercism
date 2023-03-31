package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}

	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleLayerCount, sauceLayerCount := 0, 0.0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodleLayerCount++
		case "sauce":
			sauceLayerCount++
		}
	}

	noodleRequired := noodleLayerCount * 50
	sauceRequired := sauceLayerCount * 0.2

	return noodleRequired, sauceRequired
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendRecipe, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendRecipe[len(friendRecipe)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numOfPortions int) []float64 {
	newQuantities := []float64{}

	factor := float64(numOfPortions) / 2.

	for _, quantity := range quantities {
		newQuantities = append(newQuantities, quantity*factor)
	}

	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
