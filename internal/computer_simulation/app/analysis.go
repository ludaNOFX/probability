package app

import (
	"fmt"
	"math"
	"sort"
)

type Stats struct {
	Mean     float64
	Median   float64
	Mode     float64
	Variance float64
	StdDev   float64
}

func (s Stats) String() string {
	return fmt.Sprintf(
		"Mean: %.2f, Median: %.2f, Mode: %.2f, Variance: %.2f, StdDev: %.2f",
		s.Mean, s.Median, s.Mode, s.Variance, s.StdDev,
	)
}

type Interval struct {
	Lower float64
	Upper float64
	Count int
}

func (i Interval) String() string {
	return fmt.Sprintf(
		"Lower: %.2f, Upper: %.2f, Count: %d",
		i.Lower, i.Upper, i.Count,
	)
}

func ComputeStats(data []float64) Stats {
	n := len(data)
	if n == 0 {
		return Stats{}
	}
	sum := 0.0
	min := data[0]
	max := data[0]
	freq := make(map[float64]int, n)
	for _, v := range data {
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		freq[v]++
	}
	mean := sum / float64(n)
	sorted := make([]float64, n)
	copy(sorted, data)
	sort.Float64s(sorted)
	var median float64
	if n%2 == 0 {
		median = (sorted[n/2-1] + sorted[n/2]) / 2
	} else {
		median = sorted[n/2]
	}
	var mode float64
	maxFreq := 0
	for k, v := range freq {
		if v > maxFreq {
			maxFreq = v
			mode = k
		}
	}
	varianceSum := 0.0
	for _, v := range data {
		varianceSum += (v - mean) * (v - mean)
	}
	variance := varianceSum / float64(n-1)
	stdDev := math.Sqrt(variance)
	return Stats{
		Mean:     mean,
		Median:   median,
		Mode:     mode,
		Variance: variance,
		StdDev:   stdDev,
	}
}

func Histogram(data []float64, q int) []Interval {
	if len(data) == 0 || q <= 0 {
		return nil
	}
	min := data[0]
	max := data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	r := max - min
	delta := r / float64(q)
	intervals := make([]Interval, q)
	for i := 0; i < q; i++ {
		intervals[i].Lower = min + float64(i)*delta
		intervals[i].Upper = min + float64(i+1)*delta
	}
	for _, v := range data {
		idx := int(math.Min(float64(q-1), math.Floor((v-min)/delta)))
		intervals[idx].Count++
	}
	return intervals
}
