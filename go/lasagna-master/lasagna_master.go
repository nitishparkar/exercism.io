package lasagna

const (
	defaultPreparationTimePerLayer = 2
	noodlesQuantityPerLayer        = 50
	sauceQuantityPerLayer          = 0.2
	defaultPortionsSize            = 2
)

func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = defaultPreparationTimePerLayer
	}

	return len(layers) * averagePreparationTime
}

func Quantities(layers []string) (int, float64) {
	noodlesQuantity := 0
	sauceQuantity := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodlesQuantity += noodlesQuantityPerLayer
		}

		if layer == "sauce" {
			sauceQuantity += sauceQuantityPerLayer
		}
	}

	return noodlesQuantity, sauceQuantity
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	scaledQuantities = make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] / defaultPortionsSize * float64(portions)
	}

	return
}
