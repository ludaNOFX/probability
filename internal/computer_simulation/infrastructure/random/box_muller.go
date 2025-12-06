package random

import (
	"math"
	"math/rand"
	"time"

	"github.com/ludaNOFX/probability/internal/computer_simulation/domain"
)

var _ domain.NormalGenerator = (*BoxMuller)(nil)

type BoxMuller struct {
	hasSpare bool
	spare    float64
	rng      *rand.Rand
}

func NewBoxMuller(seed int64) *BoxMuller {
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	return &BoxMuller{
		rng: rand.New(rand.NewSource(seed)),
	}
}

func (b *BoxMuller) Next() float64 {
	if b.hasSpare {
		b.hasSpare = false
		return b.spare
	}
	u1 := b.rng.Float64()
	u2 := b.rng.Float64()
	r := math.Sqrt(-2.0 * math.Log(u1))
	theta := 2.0 * math.Pi * u2
	z0 := r * math.Cos(theta)
	z1 := r * math.Sin(theta)
	b.spare = z1
	b.hasSpare = true
	return z0
}
