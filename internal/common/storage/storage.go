package storage

import "io"

type Storage interface {
	Save(path string, data io.Reader) error
}
