package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressBMP(t *testing.T) {
	inputFilename := "../test_files/image.bmp"
	outputFilename := "../test_files/outputs/compressed.bmp"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid input file", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.BMPOptions{OutputFilename: outputFilename, Depth: 24}
		err := functions.CompressBMP("non-existent-file.bmp", options)
		if err == nil {
			t.Error("Expected error for non-existent input file")
		}
	})

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.BMPOptions{OutputFilename: "", Depth: 24}
		err := functions.CompressBMP(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("invalid depth option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.BMPOptions{OutputFilename: outputFilename, Depth: 0}
		err := functions.CompressBMP(inputFilename, options)

		if err == nil {
			t.Fatalf("Expected error for invalid depth option")
		}

		options = functions.BMPOptions{OutputFilename: outputFilename, Depth: 64}
		err = functions.CompressBMP(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid depth option")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.BMPOptions{OutputFilename: outputFilename, Depth: 24}
		err := functions.CompressBMP(inputFilename, options)
		if err != nil {
			t.Fatalf("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() > inputFileInfo.Size() {
			t.Error("BMP compression did not reduce the file size")
		}
	})
}
