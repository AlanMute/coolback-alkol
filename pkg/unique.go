package pkg

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetPath(name string, folder string) string {
	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(name, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	dirPath := filepath.Join(folder, uniqueName)

	return dirPath
}

func CreateUniqueFolder(name string, folder string) (string, error) {
	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(name, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	dirPath := filepath.Join(folder, uniqueName)

	splitFileName := strings.Split(dirPath, "\\")
	dbFolderName := splitFileName[len(splitFileName)-1]

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
		return dbFolderName, nil
	}

	return "", fmt.Errorf("directory already exists")
}

func CreateUniqueFile(file multipart.File, name string, folder string, requiredExt string) (string, error) {
	ext := filepath.Ext(name)
	if ext != requiredExt {
		return "", fmt.Errorf("wrong file extension")
	}

	nameWithoutExt := name[:len(name)-len(ext)]

	validName := regexp.MustCompile(`[^\p{L}0-9\s-]`).ReplaceAllString(nameWithoutExt, "")

	words := strings.Fields(validName)
	uniqueName := strings.Join(words, "-")

	filePath := filepath.Join(folder, uniqueName+ext)

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		out, err := os.Create(filePath)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(out, file)
		if err != nil {
			return "", err
		}

		if err = out.Close(); err != nil {
			return "", err
		}

		return uniqueName + ext, nil
	}

	if err = file.Close(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("file already exists")
}
