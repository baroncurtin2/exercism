package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutesPerLayer int) int {
	multiplier := func(minutesPerLayer int) int {
		if minutesPerLayer == 0 {
			return 2
		}
		return minutesPerLayer
	}(minutesPerLayer)

	return len(layers) * multiplier
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodlesQuantity int, sauceQuantity float64) {
	const noodlesGramsPerLayer int = 50
	const sauceLitersPerLayer float64 = 0.2

	noodleCount := 0
	sauceCount := 0

	for _, v := range layers {
		if v == "noodles" {
			noodleCount++
		}

		if v == "sauce" {
			sauceCount++
		}
	}

	return noodleCount * noodlesGramsPerLayer, float64(sauceCount) * sauceLitersPerLayer
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	nf := len(friendsList)
	nm := len(myList)

	myList[nm-1] = friendsList[nf-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	amounts := make([]float64, len(quantities))

	for i, v := range quantities {
		amounts[i] = v * float64(portions) / 2
	}
	return amounts
}
