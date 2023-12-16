package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, v := range layers {
		if v == "sauce" {
			sauce += 1
		} else if v == "noodles" {
			noodles += 1
		}
	}
	return noodles * 50, sauce * 0.2
}

func AddSecretIngredient(friendsLayers, layers []string) {
	layers[len(layers)-1] = friendsLayers[len(friendsLayers)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := []float64{}
	for _, v := range quantities {
		scaled = append(scaled, v*float64(portions)/2)
	}
	return scaled
}
