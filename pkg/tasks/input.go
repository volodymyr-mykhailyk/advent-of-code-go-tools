package tasks

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func PathToFile(file string) string {
	basePath := os.Getenv("DATA_DIR")
	if basePath == "" {
		ex, _ := os.Getwd()
		basePath = filepath.Dir(ex)
	}
	return filepath.Join(basePath, file)
}

func ReadLines(path string) []string {
	body, err := os.ReadFile(PathToFile(path))
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return strings.Split(string(body), "\n")
}

func StreamLines(path string, lines chan<- string) {
	file, err := os.Open(PathToFile(path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer close(lines)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	lines <- scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Iterate(inputs []string, iterator func(line string, i int)) {
	for i, line := range inputs {
		iterator(line, i)
	}
}
