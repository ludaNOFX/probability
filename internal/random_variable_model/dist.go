package randomvariablemodel

import (
	"math"

	"gonum.org/v1/plot/plotter"
)

type Params struct {
	A float64
	C float64
}

func NewParams(a float64, f func(a float64) float64) Params {
	return Params{A: a, C: f(a)}
}

func CalcC(a float64) float64 {
	return 1.0 / (0.5 + a)
}

func Fx(x float64, p Params) float64 {
	if x >= 0 && x <= 1 {
		return p.C * (x + p.A)
	}
	return 0
}

func F(x float64, p Params) float64 {
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return p.C * (0.5*x*x + p.A*x)
}

func GeneratePoints(
	p Params,
	xmin, xmax, step float64,
	f func(x float64, p Params) float64,
) plotter.XYs {
	n := int((xmax-xmin)/step) + 2
	pts := make(plotter.XYs, 0, n)

	for x := xmin; x <= xmax; x += step {
		pts = append(pts, plotter.XY{
			X: x,
			Y: f(x, p),
		})
	}
	return pts
}

func Mean(p Params) float64 {
	return p.C * (1.0/3.0 + p.A/2.0)
}

func Mean2(p Params) float64 {
	return p.C * (1.0/4.0 + p.A/3.0)
}

func Variance(p Params) float64 {
	m1 := Mean(p)
	m2 := Mean2(p)
	return m2 - m1*m1
}

func Sigma(p Params) float64 {
	return math.Sqrt(Variance(p))
}

func Mode(p Params) float64 {
	if p.C > 0 {
		return 1.0
	} else if p.C < 0 {
		return 0.0
	} else {
		return 0.5
	}
}

func Median(p Params) float64 {
	lo := 0.0
	hi := 1.0

	for i := 0; i < 80; i++ {
		m := (lo + hi) / 2.0
		if F(m, p) < 0.5 {
			lo = m
		} else {
			hi = m
		}
	}

	return (lo + hi) / 2.0
}
