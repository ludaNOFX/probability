package statistics

func Mean(xs []float64) (float64, error) {
	n := len(xs)
	if n == 0 {
		return 0, emptyError
	}
	var total float64
	for _, x := range xs {
		total += x
	}
	return total / float64(n), nil
}
