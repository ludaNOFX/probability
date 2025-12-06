package app

import "math/rand"

func UniformSample(a, b float64, n int, seed int64) []float64 {
	if seed == 0 {
		seed = rand.Int63()
	}
	rng := rand.New(rand.NewSource(seed))
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		out[i] = a + (b-a)*rng.Float64()
	}
	return out
}
