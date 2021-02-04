package services

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"mime/multipart"
	"path/filepath"
	"sweet_fantasy_go/internal/models"
)

func CreateAndSaveFile(file *multipart.FileHeader, filePath string) (*models.File, error) {
	relativePath := getRelativePath(file.Filename, filePath)

	fullPath, err := getFullPath(relativePath)
	if err != nil {
		return nil, err
	}

	err = fasthttp.SaveMultipartFile(file, fullPath)
	if err != nil {
		return nil, err
	}

	return &models.File{Path: relativePath}, nil
}

func getRelativePath(fileName string, filePath string) string {
	return fmt.Sprintf(
		"%s%s%s",
		filePath,
		string(filepath.Separator),
		fileName,
	)
}

func getFullPath(relativePath string) (string, error) {
	return filepath.Abs(fmt.Sprintf(
		"../assets/%s",
		relativePath,
	))
}
