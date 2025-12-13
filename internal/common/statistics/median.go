package statistics

func Median(xs []float64) (float64, error) {
	return Quantile(xs, 0.5)
}
