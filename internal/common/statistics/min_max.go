package statistics

func Min(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, emptyError
	}
	minElem := xs[0]
	for _, elem := range xs {
		if elem < minElem {
			minElem = elem
		}
	}
	return minElem, nil
}

func Max(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, emptyError
	}
	maxElem := xs[0]
	for _, elem := range xs {
		if elem > maxElem {
			maxElem = elem
		}
	}
	return maxElem, nil
}
