package sample

type Generator func() float64

func Sample(n int, gen Generator) []float64 {
	result := make([]float64, n)
	for i := 0; i < n; i++ {
		result[i] = gen()
	}
	return result
}
