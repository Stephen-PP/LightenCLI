package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressJPEG(t *testing.T) {
	inputFilename := "../test_files/image.jpg"
	outputFilename := "../test_files/outputs/compressed_jpeg.jpg"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.JpegOptions{Quality: 80, OutputFilename: ""}
		err := functions.CompressJPEG(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("invalid quality level", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.JpegOptions{Quality: 0, OutputFilename: outputFilename}
		err := functions.CompressJPEG(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid quality level")
		}
	})

	t.Run("compression with default options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.JpegOptions{Quality: 80, OutputFilename: outputFilename}
		err := functions.CompressJPEG(inputFilename, options)
		if err != nil {
			t.Error("JPEG compression failed:", err)
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("JPEG compression did not reduce the file size")
		}
	})
}
