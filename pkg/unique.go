package pkg

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func UniqueFolder(name string, folder string) (string, error) {
	for {
		uniqueID := uuid.New().String()
		uniqueName := name + "_" + uniqueID
		filePath := filepath.Join(folder, uniqueName)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return filePath, nil
		}
	}
}

func UniqueFile(ext string, folder string) (string, error) {
	for {
		uniqueID := uuid.New().String()
		uniqueName := "lesson_" + uniqueID + ext
		filePath := filepath.Join(folder, uniqueName)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return filePath, nil
		}
	}
}
