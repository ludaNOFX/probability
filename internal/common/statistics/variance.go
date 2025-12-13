package statistics

func Variance(xs []float64) (float64, error) {
	mean, err := Mean(xs)
	if err != nil {
		return 0, err
	}
	var sumSquares float64
	for _, x := range xs {
		diff := x - mean
		sumSquares += diff * diff
	}
	return sumSquares / float64(len(xs)), nil
}
