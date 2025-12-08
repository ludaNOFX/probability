package random

import (
	"math/rand"
	"time"

	"github.com/ludaNOFX/probability/internal/computer_simulation/domain"
)

var _ domain.UniformGenerator = (*Uniform)(nil)

// Uniform реализует генератор равномерных чисел U(0,1)
type Uniform struct {
	rng *rand.Rand
}

func NewUniform(seed int64) *Uniform {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	return &Uniform{
		rng: rand.New(rand.NewSource(seed)),
	}
}

// Float64 возвращает U в (0,1)
func (u *Uniform) Float64() float64 {
	return u.rng.Float64()
}
