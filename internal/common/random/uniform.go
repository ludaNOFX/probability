package random

func (s *Source) Uniform01() float64 {
	return s.rng.Float64()
}

func (s *Source) Uniform(a, b float64) float64 {
	return a + (b-a)*s.Uniform01()
}
