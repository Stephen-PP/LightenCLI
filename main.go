package main

import (
	"fmt"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func main() {
	mp3Options := functions.MP3Options{Bitrate: 64, Quality: 9, OutputFilename: "./test_files/outputs/compressed_mp3.mp3"}
	err := functions.CompressMP3("./test_files/sample.mp3", mp3Options)
	if err != nil {
		fmt.Println(err)
	}
}
