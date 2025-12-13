package random

import (
	"math/rand"
	"time"
)

type Source struct {
	rng      *rand.Rand
	hasSpare bool
	spare    float64
}

func NewSource(seed int64) *Source {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	return &Source{
		rng: rand.New(rand.NewSource(seed)),
	}
}
