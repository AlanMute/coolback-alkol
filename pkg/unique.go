package pkg

import (
	"os"
	"path/filepath"
	"strconv"
)

func UniqueFolder(name string, folder string) (string, error) {
	for i := 1; ; i++ {
		uniqueName := name + "_" + strconv.Itoa(i)
		filePath := filepath.Join(folder, uniqueName)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return folder + "/" + uniqueName, nil
		} else {
			return "", err
		}
	}
}

func UniqueFile(folder string) (string, error) {
	name := "lesson.md"

	extension := filepath.Ext(name)
	nameWithoutExt := name[:len(name)-len(extension)]

	for i := 1; ; i++ {
		uniqueName := nameWithoutExt + "_" + strconv.Itoa(i) + extension
		filePath := filepath.Join(folder, uniqueName)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return folder + "/" + uniqueName, nil
		} else {
			return "", err
		}
	}
}
