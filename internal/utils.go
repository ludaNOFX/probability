package internal

import (
	"os"
	"path/filepath"
)

const BaseStorage string = "filestore"

type PrjMap struct {
	prjName string
	types   []string
}

func NewPrjMap(prjName string, storageTypes []string) *PrjMap {
	return &PrjMap{prjName: prjName, types: storageTypes}
}

type PathMap map[string]string

func findStartDir() (string, error) {
	startDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return startDir, err
}

// The nesting parameter specifies how far from the parent directory the process is launched
// For example, if the process starts /home/cmd/main.go
// and you want to find home as root dir, then the directory is nested by 1.
func FindRootDir(nesting int) (string, error) {
	startDir, err := findStartDir()
	if err != nil {
		return "", err
	}
	if nesting < 1 {
		return startDir, nil
	}
	rootDir := filepath.Dir(startDir)
	for i := 1; i < nesting; i++ {
		rootDir = filepath.Dir(rootDir)
	}
	return rootDir, nil
}

func BuildStoragePathOne(rootDir, prjName, storageType string) string {
	return filepath.Join(rootDir, BaseStorage, prjName, storageType)
}

func BuildStoragePathMany(rootDir string, project *PrjMap) PathMap {
	pathMap := make(PathMap, len(project.types))
	for _, t := range project.types {
		pathMap[t] = BuildStoragePathOne(
			rootDir,
			project.prjName,
			t,
		)
	}
	return pathMap
}
