package pkg

import (
	"bufio"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
)

func GetPathToFile(name string, ext string, folder string) (string, error) {
	validName := regexp.MustCompile(`[^\d*[1-9]\d*`).ReplaceAllString(name, "")

	dirPath := filepath.Join(folder, validName+ext)

	_, err := os.Stat(dirPath)
	if !os.IsNotExist(err) {
		return dirPath, nil
	}

	return "", fmt.Errorf("folder with that name doesn't exist")
}

func GenerateUniqueFile(fileName string, fileID string, folder string, requiredExt string) (string, error) {
	ext := filepath.Ext(fileName)
	if ext != requiredExt {
		return "", fmt.Errorf("wrong file extension")
	}

	nameWithoutExt := fileID

	validName := regexp.MustCompile(`[^\d*[1-9]\d*`).ReplaceAllString(nameWithoutExt, "")

	filePath := filepath.Join(folder, validName+ext)

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return filePath, nil
	}

	return "", fmt.Errorf("file already exists")
}

func CreateFile(file multipart.File, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	if err = out.Close(); err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	return nil
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Error opening a file: " + err.Error())
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error scanning a file: " + err.Error())
	}

	return lines, nil
}
