package services

import (
	"crypto/sha1"
	"fmt"
	"github.com/valyala/fasthttp"
	"mime/multipart"
	"path/filepath"
	"strings"
	db "sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/models"
	"sweet_fantasy_go/internal/repositories/files_repository"
	"time"
)

func CreateAndSaveFile(file *multipart.FileHeader, filePath string) (*models.File, error) {
	relativePath, err := saveFile(file, filePath)
	if err != nil {
		return nil, err
	}
	fileModel := &models.File{Path: relativePath}
	files_repository.Create(fileModel)
	return fileModel, nil
}

func ReplaceFile(newFile *multipart.FileHeader, filePath string, prevFile *models.File) error {
	relativePath, err := saveFile(newFile, filePath)
	if err != nil {
		return err
	}
	if err = prevFile.RemoveFromStorage(); err != nil {
		return err
	}
	prevFile.Path = relativePath
	db.DBConn.Save(&prevFile)
	return nil
}

func saveFile(file *multipart.FileHeader, filePath string) (relative string, err error) {
	relativePath := getRelativePath(generateFileName(file), filePath)
	fullPath, err := getFullPath(relativePath)
	if err != nil {
		return "", err
	}

	err = fasthttp.SaveMultipartFile(file, fullPath)
	if err != nil {
		return "", err
	}
	return relativePath, err
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

func generateFileName(file *multipart.FileHeader) string {
	separatedString := strings.Split(file.Filename, ".")
	return strings.Join([]string{generateTimestampHash(), separatedString[len(separatedString)-1]}, ".")
}

func generateTimestampHash() string {
	timestampBytes := []byte(time.Now().Format("2006-01-02 15:04:05"))
	return fmt.Sprintf("%x", sha1.Sum(timestampBytes))
}
