package statistics

import "math"

func StdDev(xs []float64) (float64, error) {
	variance, err := Variance(xs)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(variance), nil
}
