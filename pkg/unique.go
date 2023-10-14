package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func CreateFile(path string, file []string) error {
	content := strings.Join(file, "\n")

	err := os.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFile(path string, file []string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("File does not exist for path: %s", path)
	} else if err != nil {
		return err
	}

	content := strings.Join(file, "\n")

	err = os.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
