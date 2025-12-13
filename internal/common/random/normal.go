package random

import "math"

func (s *Source) Normal01() float64 {
	return s.normalBoxMuller()
}

func (s *Source) Normal(mu, sigma float64) float64 {
	return mu + sigma*s.Normal01()
}

func (s *Source) normalBoxMuller() float64 {
	if s.hasSpare {
		s.hasSpare = false
		return s.spare
	}
	u1 := s.Uniform01()
	u2 := s.Uniform01()
	r := math.Sqrt(-2.0 * math.Log(u1))
	theta := 2.0 * math.Pi * u2
	z0 := r * math.Cos(theta)
	z1 := r * math.Sin(theta)
	s.spare = z1
	s.hasSpare = true
	return z0
}
