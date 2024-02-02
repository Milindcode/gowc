package main

import (
	"fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	const demo_file_name = "test.txt"

	data, err := os.ReadFile(demo_file_name)
	check(err)
	fmt.Print(len(string(data)))
	fmt.Print(os.Args[2])
}