package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func LoadString(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Error loading file '%s': %s\n", filename, err))
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, string(scanner.Text()))
	}

	return lines
}
