package impl

import (
	"fmt"
	"io"

	"github.com/ludaNOFX/probability/internal/common/plt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	msgFilureHist string = "Building histogram failure:"
	msgFilureLine string = "Building line plot failure:"
)

type GonumPlotter struct{}

func (gP *GonumPlotter) configureBase(p *plot.Plot, cfg plt.Meta) {
	p.Title.Text = cfg.Title
	p.X.Label.Text = cfg.XLabel
	p.Y.Label.Text = cfg.YLabel
}

func (gP *GonumPlotter) configureLine(l *plotter.Line, cfg plt.LineVizulizer) {
	l.Color = cfg.Color
	l.Width = vg.Length(cfg.Width)
}

func (gP *GonumPlotter) configureHist(h *plotter.Histogram, cfg plt.LineVizulizer) {
	h.Color = cfg.Color
	h.Width = cfg.Width
}

func (gP *GonumPlotter) toPlotterType(xys plt.PlotData) plotter.XYs {
	dst := make(plotter.XYs, len(xys))
	for idx := range xys {
		dst[idx].X = xys[idx].X
		dst[idx].Y = xys[idx].Y
	}
	return dst
}

func (gP *GonumPlotter) defineMaxY(xys plt.PlotData) float64 {
	maxY := 0.0
	for _, pt := range xys {
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}
	return maxY
}

func (gP *GonumPlotter) buildVL(item plt.VerticalLine, maxY float64) (*plotter.Line, error) {
	vl, err := plotter.NewLine(
		plotter.XYs{
			{X: item.X, Y: 0},
			{X: item.X, Y: maxY},
		},
	)
	if err != nil {
		return nil, err
	}
	gP.configureLine(vl, item.LV)
	return vl, nil
}

func (gP *GonumPlotter) buildLableVL(item plt.VerticalLine, maxY float64) (*plotter.Labels, error) {
	const fivePercent float64 = 0.95
	label, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    []plotter.XY{{X: item.X, Y: maxY * fivePercent}},
		Labels: []string{item.Label},
	})
	if err != nil {
		return nil, err
	}
	return label, nil
}

func (gP *GonumPlotter) LinePlot(xys plt.PlotData, cfg plt.DistributionConfig, out io.Writer) error {
	p := plot.New()
	gP.configureBase(p, cfg.Info.Meta)
	line, err := plotter.NewLine(gP.toPlotterType(xys))
	if err != nil {
		return fmt.Errorf("%s %w", msgFilureLine, err)
	}
	gP.configureLine(line, cfg.LV)
	p.Add(line)
	maxY := gP.defineMaxY(xys)
	for _, item := range cfg.VLines {
		vl, err := gP.buildVL(item, maxY)
		if err != nil {
			return fmt.Errorf("%s %w", msgFilureLine, err)
		}
		p.Add(vl)
		label, err := gP.buildLableVL(item, maxY)
		if err != nil {
			return fmt.Errorf("%s %w", msgFilureLine, err)
		}
		p.Add(label)

	}
	wt, err := gP.configureFormat(p, cfg.F)
	if err != nil {
		return fmt.Errorf("%s %w", msgFilureLine, err)
	}
	if _, err := wt.WriteTo(out); err != nil {
		return fmt.Errorf("%s %w", msgFilureLine, err)
	}
	return nil
}

func (gP *GonumPlotter) configureFormat(p *plot.Plot, cfg plt.Format) (io.WriterTo, error) {
	wt, err := p.WriterTo(vg.Points(cfg.W), vg.Points(cfg.H), cfg.Ext)
	if err != nil {
		return nil, err
	}
	return wt, nil
}

func (gP *GonumPlotter) Histogram(data plt.HistData, cfg plt.HistogramConfig, out io.Writer) error {
	p := plot.New()
	gP.configureBase(p, cfg.Info.Meta)
	values := plotter.Values(data)
	h, err := plotter.NewHist(values, cfg.Bins)
	if err != nil {
		return fmt.Errorf("%s %w", msgFilureHist, err)
	}
	gP.configureHist(h, cfg.Info.LV)
	p.Add(h)
	wt, err := gP.configureFormat(p, cfg.F)
	if err != nil {
		return fmt.Errorf("%s %w", msgFilureHist, err)
	}
	if _, err := wt.WriteTo(out); err != nil {
		return fmt.Errorf("%s %w", msgFilureHist, err)
	}
	return nil
}
