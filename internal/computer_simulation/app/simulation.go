package app

import "github.com/ludaNOFX/probability/internal/computer_simulation/domain"

type NormalSimulation struct {
	generator domain.NormalGenerator
}

func NewNormalGenerator(g domain.NormalGenerator) *NormalSimulation {
	return &NormalSimulation{
		generator: g,
	}
}

func (s *NormalSimulation) Sample(n int) []float64 {
	out := make([]float64, n)
	for i := 0; i < n; i++ {
		out[i] = s.generator.Next()
	}
	return out
}
