package utils

import (
	"bufio"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Line( path string ) int {

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		count ++
	}

	if err := scanner.Err(); err != nil {
        log.Fatalf("error while scanning file: %s", err)
    }

	return count
}