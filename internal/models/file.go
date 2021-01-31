package models

import (
	"os"
	"path/filepath"
)

type File struct {
	Base
	Path string
}

func (file *File) removeFromStorage() error {
	path, err := filepath.Abs(file.Path)
	if err != nil {
		return err
	}
	return os.Remove(path)
}
