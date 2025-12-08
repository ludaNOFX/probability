package randomvariablemodel

import (
	"image/color"
	"os"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Plot(
	pts plotter.XYs,
	filename string,
	meta *MetaInfo,
	lines ...float64,
) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	p := plot.New()
	p.Title.Text = meta.Title
	p.X.Label.Text = meta.XLabel
	p.Y.Label.Text = meta.YLabel

	line, err := plotter.NewLine(pts)
	if err != nil {
		return err
	}
	line.Color = color.Black
	line.Width = 2.0
	p.Add(line)

	maxY := 0.0
	for _, pt := range pts {
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}

	colors := []color.Color{
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{0, 255, 0, 255},
		color.RGBA{255, 165, 0, 255},
	}

	labels := []string{"M", "Me", "Mo"}

	for i, x := range lines {
		if i >= len(colors) {
			break
		}

		vl, err := plotter.NewLine(plotter.XYs{
			{X: x, Y: 0},
			{X: x, Y: maxY},
		})
		if err != nil {
			return err
		}
		vl.Color = colors[i]
		vl.Width = 1.5
		p.Add(vl)

		if i < len(labels) {
			label, err := plotter.NewLabels(plotter.XYLabels{
				XYs:    []plotter.XY{{X: x, Y: maxY * 0.95}},
				Labels: []string{labels[i]},
			})
			if err != nil {
				return err
			}
			p.Add(label)
		}
	}

	if err := p.Save(vg.Length(meta.W), vg.Length(meta.H), filename); err != nil {
		return err
	}
	return nil
}
