package models

import (
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type File struct {
	gorm.Model
	Path string
}

func (file *File) removeFromStorage() error {
	path, err := filepath.Abs(file.Path)
	if err != nil {
		return err
	}
	return os.Remove(path)
}
