package main

import (
	"fmt"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func main() {
	// Lossless
	losslessOpts := functions.LosslessPngOptions{Level: 6, Strip: true, OutputFilename: "./test_files/outputs/image-lossless.png"}
	err := functions.CompressPNGLossless("./test_files/image.png", losslessOpts)
	if err != nil {
		fmt.Println(err)
	}

	// Lossy
	lossyOpts := functions.LossyPngOptions{Quality: "65-80", Speed: 5, OutputFilename: "./test_files/outputs/image-lossy.png"}
	err = functions.CompressPNGLossy("./test_files/image.png", lossyOpts)
	if err != nil {
		fmt.Println(err)
	}
}
