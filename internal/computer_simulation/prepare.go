package computersimulation

import (
	"fmt"
	"os"

	"github.com/ludaNOFX/probability/internal"
)

func SetupPaths(project *internal.PrjMap) (internal.PathMap, error) {
	rootDir, err := internal.FindRootDir(2)
	if err != nil {
		return nil, fmt.Errorf("Setup path error: %w", err)
	}
	pathMap := internal.BuildStoragePathMany(rootDir, project)
	for _, path := range pathMap {
		if err := os.MkdirAll(path, 0o755); err != nil {
			return nil, fmt.Errorf("Setup %s error: %w", path, err)
		}
	}
	return pathMap, nil
}
