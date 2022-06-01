package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	filename := "./yinhun.mp4"
	width := 640
	height := 360
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprint("%dx%d", width, height), "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		panic("could not generate frame")
	}

}
