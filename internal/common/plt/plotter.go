package plt

import (
	"io"
)

type Plotter interface {
	// TODO нужно будет еще подумать что может получать и возвращать интерфейс
	LinePlot(xys PlotData, cfg DistributionConfig, out io.Writer) error
	// TODO нужно будет еще подумать что может получать и возвращать интерфейс
	Histogram(data HistData, cfg HistogramConfig, out io.Writer) error
}
