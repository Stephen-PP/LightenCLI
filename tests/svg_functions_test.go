package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressSVG(t *testing.T) {
	inputFilename := "../test_files/image.svg"
	outputFilename := "../test_files/outputs/compressed.svg"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid input file", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.SvgOptions{OutputFilename: outputFilename}
		err := functions.CompressSVG("non-existent-file.svg", options)
		if err == nil {
			t.Error("Expected error for non-existent input file")
		}
	})

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.SvgOptions{OutputFilename: ""}
		err := functions.CompressSVG(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.SvgOptions{OutputFilename: outputFilename}
		err := functions.CompressSVG(inputFilename, options)
		if err != nil {
			t.Error("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("SVG compression did not reduce the file size")
		}
	})
}
