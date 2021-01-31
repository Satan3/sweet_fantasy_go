package models

import (
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Base
	Path string `json:"path"`
}

func (file *File) removeFromStorage() error {
	relativePath := strings.Join([]string{"../assets/", file.Path}, "")
	path, err := filepath.Abs(relativePath)
	if err != nil {
		return err
	}
	return os.Remove(path)
}
