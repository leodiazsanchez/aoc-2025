package utils

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	lines := make([]string, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
