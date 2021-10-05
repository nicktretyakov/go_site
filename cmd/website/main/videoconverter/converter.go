package main

import (
	//"bufio"
	"fmt"
	"os"

	//"os/exec"
	"path/filepath"
)

func converter() {
	fmt.Println("starting transcoder\n")
	// getting the directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	} else {
		// calling the transcoder
		result := transcoder("test.mp4", "logo.png", dir)
		if result {
			fmt.Println("Success")
		} else {
			fmt.Println("Error")
		}

	}

}
