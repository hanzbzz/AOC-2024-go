package utils

import (
	"os"
	"strings"
)

func ReadFileLines(path string) ([]string, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fileString := string(fileBytes)
	return strings.Split(fileString, "\n"), nil
}
