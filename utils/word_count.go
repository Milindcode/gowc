package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)
func Word( path string ) int {

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
        words := strings.Fields(line)
		count += len(words)
	}

	if err := scanner.Err(); err != nil {
        log.Fatalf("error while scanning file: %s", err)
    }

	return count
}