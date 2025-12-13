package impl

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

type FileStorage struct {
	baseDir string
}

func NewFileStorage(baseDir string) *FileStorage {
	return &FileStorage{baseDir: baseDir}
}

func (s *FileStorage) preparePath(path string) (string, error) {
	filePath := filepath.Join(s.baseDir, path)
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (s *FileStorage) Save(path string, data io.Reader) error {
	fullPath, err := s.preparePath(path)
	if err != nil {
		return err
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = io.Copy(w, data)
	if err != nil {
		return err
	}
	return w.Flush()
}
