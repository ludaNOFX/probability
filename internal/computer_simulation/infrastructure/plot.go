package infrastructure

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PlotHistogram(data []float64, filename string, bins int) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	p := plot.New()
	base := filepath.Base(filename)
	baseWithoutExt, ok := strings.CutSuffix(base, filepath.Ext(base))
	if !ok {
		return errors.New("file base not have ext")
	}
	p.Title.Text = baseWithoutExt
	p.X.Label.Text = "Value"
	p.Y.Label.Text = "Frequency"

	values := make(plotter.Values, len(data))
	copy(values, data)

	h, err := plotter.NewHist(values, bins)
	if err != nil {
		return err
	}

	p.Add(h)

	return p.Save(10*vg.Centimeter, 10*vg.Centimeter, filename)
}
