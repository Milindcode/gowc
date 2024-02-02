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
	var demo_file_name string= "test.txt"

	dir_path, err := os.Getwd()
	check(err)

	operation_choice := ""
	if( len(os.Args) >= 2){
		operation_choice = os.Args[1]
	}

	if( len(os.Args) == 3){
		demo_file_name = dir_path + "\\" + os.Args[2]
	}

	data, err := os.ReadFile(demo_file_name)
    check(err)

	switch operation_choice {
	case "-c":
		fmt.Println(character(string(data)))
	case "-w":
		fmt.Println("Word Count")
	case "-l":
		fmt.Println("Line Count")
	case "-m":
		fmt.Println("Somthing Count")
	default:
		fmt.Println("Default")
	}

}