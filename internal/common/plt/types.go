package plt

import "image/color"

// Типы которые используются с интерфейсом по сути иметь будут валидацию конструкторы итд
type Point struct {
	X, Y float64
}

type PlotData []Point

type HistData []float64

type Meta struct {
	Title  string
	XLabel string
	YLabel string
}

type Format struct {
	W, H float64
	Ext  string // TODO НУЖНО БУДЕТ ВАЛИДИРОВАТЬ
}

type Info struct {
	Meta Meta
	LV   LineVizulizer
	F    Format
}

type LineVizulizer struct {
	Width float64
	Color color.Color
}

type VerticalLine struct {
	X     float64
	LV    LineVizulizer
	Label string
}

type DistributionConfig struct {
	VLines []VerticalLine
	Info
}

type HistogramConfig struct {
	Info
	Bins int
}

//
