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
	var demo_file_name string= "data/test.txt"

	dir_path, err := os.Getwd()
	check(err)

	operation_choice := ""
	if( len(os.Args) >= 2){
		if( os.Args[1][0] == '-' ){
			operation_choice = os.Args[1]
		}else {
			demo_file_name = dir_path + "\\" + os.Args[1]
		}
	}

	if( len(os.Args) == 3){
		demo_file_name = dir_path + "\\" + os.Args[2]
	}

	switch operation_choice {
	case "-c":
		fmt.Println(character(demo_file_name))
	case "-w":
		fmt.Println(word(demo_file_name))
	case "-l":
		fmt.Println(line(demo_file_name))
	case "-m":
		fmt.Println(multibyte(demo_file_name))
	default:
		fmt.Println(line(demo_file_name), word(demo_file_name), character(demo_file_name), os.Args[1])
	}

}