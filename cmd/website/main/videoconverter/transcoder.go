package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

func transcoder(file, logo, dir string) bool {

	fmt.Printf("converting to 144\n")
	makeWithRes(144, 256, file, logo, dir)
	fmt.Printf("converting to 240\n")
	makeWithRes(240, 428, file, logo, dir)
	fmt.Printf("converting to 360\n")
	makeWithRes(360, 640, file, logo, dir)
	fmt.Printf("converting to 480\n")
	makeWithRes(480, 854, file, logo, dir)
	fmt.Printf("converting to 720\n")
	makeWithRes(720, 1280, file, logo, dir)
	fmt.Printf("converting to 1080\n")
	makeWithRes(1080, 1920, file, logo, dir)

	return true

}

// creates a new video with new resolution and logo on top right
func makeWithRes(height int, width int, file, logo, location string) bool {

	file_path := "workspace/" + file
	logo_file := "workspace/" + logo

	w := strconv.Itoa(width)
	h := strconv.Itoa(height)

	scale_command := "scale=" + w + ":" + h
	//logo_command := "overlay=5:5" // Top left
	logo_command := "overlay=main_w-overlay_w-5:5" //"overlay=W-w-5:5" // Top right
	//logo_command := "overlay=5:H-h-5" // Bottom left
	//logo_command := "overlay=W-w-5:H-h-5" // Bottom right
	// filter commands
	command := logo_command + "," + scale_command
	cmd := exec.Command("ffmpeg", "-i", file_path, "-i", logo_file, "-filter_complex", command, "result/"+w+"_x_"+h+"_"+file)
	cmd.Dir = location
	// Catching errors
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return false
	}
	result := out.String()
	if result != "" {
		fmt.Println("Result: " + result)
	} else {
		fmt.Println("Result: success")
	}

	return true
}
