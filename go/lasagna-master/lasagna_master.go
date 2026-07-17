package lasagnamaster

func PreparationTime(layers []string, preperationTimePerLayer int) int {
	if preperationTimePerLayer == 0 {
		preperationTimePerLayer = 2
	}

	return len(layers) * preperationTimePerLayer
}

func Quantities(layers []string) (noodles int, souce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
			continue
		}

		if layer == "sauce" {
			souce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendsList, myList []string) {
	lastIndexOfMyList := len(myList) - 1
	lastIndexOfFriendsList := len(friendsList) - 1

	myList[lastIndexOfMyList] = friendsList[lastIndexOfFriendsList]
}

// quantities := []float64{ 1.2, 3.6, 10.5 }
// scaledQuantities := ScaleRecipe(quantities, 4)
// // => []float64{ 2.4, 7.2, 21 }

func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := []float64{}

	for _, quantity := range quantities {
		onePortionQuantity := quantity / 2
		newValue := onePortionQuantity * float64(portions)
		newQuantities = append(newQuantities, newValue)
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
