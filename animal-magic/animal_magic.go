package chance

import "math/rand"

func RollADie() int {
	return rand.Intn(19) + 1
}

func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

func ShuffleAnimals() []string {
	animals := []string{
		"ant",
		"beaver",
		"cat",
		"dog",
		"elephant",
		"fox",
		"giraffe",
		"hedgehog",
	}

	rand.Shuffle(8, func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})

	return animals
}
