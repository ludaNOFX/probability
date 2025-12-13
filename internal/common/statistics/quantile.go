package statistics

import (
	"errors"
	"math"
	"slices"
)

func Quantile(xs []float64, p float64) (float64, error) {
	if len(xs) == 0 {
		return 0, errors.New("slice is empty")
	}
	if p < 0 || p > 1 {
		return 0, errors.New("p must be between 0 and 1")
	}

	sorted := make([]float64, len(xs))
	copy(sorted, xs)
	slices.Sort(sorted)
	n := len(sorted)
	if p == 0 {
		return sorted[0], nil
	}
	if p == 1 {
		return sorted[n-1], nil
	}
	pos := p * float64(n-1)
	lower := int(math.Floor(pos))
	upper := int(math.Ceil(pos))
	if upper == lower {
		return sorted[lower], nil
	}
	weight := pos - float64(lower)
	return sorted[lower]*(1-weight) + sorted[upper]*weight, nil
}
