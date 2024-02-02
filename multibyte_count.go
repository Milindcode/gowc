package main

import (
	"bufio"
	"log"
	"os"
	"unicode/utf8"
)

func multibyte( path string ) int {

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += utf8.RuneCountInString(line) + 2
	}

	if err := scanner.Err(); err != nil {
        log.Fatalf("error while scanning file: %s", err)
    }

	return count
}