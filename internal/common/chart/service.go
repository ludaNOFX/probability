package chart

import (
	"github.com/ludaNOFX/probability/internal/common/plt"
	"github.com/ludaNOFX/probability/internal/common/storage"
)

type ChartService struct {
	P plt.Plotter
	S storage.Storage
}
