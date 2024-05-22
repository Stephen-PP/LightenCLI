package main

import (
	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func main() {
	losslessOpts := functions.LosslessOptions{Level: 7, Strip: true, Interlace: true, OutputFilename: "./test_files/outputs/output-lossless.png"}
	_ = functions.CompressPNGLossless("./test_files/image.png", losslessOpts)

	// Lossy
	lossyOpts := functions.LossyOptions{Quality: "65-80", Speed: 5, OutputFilename: "./test_files/outputs/output-lossy.png"}
	_ = functions.CompressPNGLossy("./test_files/image.png", lossyOpts)
}
