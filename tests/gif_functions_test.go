package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressGifLossless(t *testing.T) {
	inputFilename := "../test_files/image.gif"
	outputFilename := "../test_files/outputs/compressed_lossless.gif"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessGifOptions{OutputFilename: ""}
		err := functions.CompressGIFLossless(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("compression with no specified options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessGifOptions{}
		err := functions.CompressGIFLossless(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessGifOptions{OutputFilename: outputFilename}
		err := functions.CompressGIFLossless(inputFilename, options)
		if err != nil {
			t.Error("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossless compression did not reduce the file size")
		}
	})
}

func TestCompressGifLossy(t *testing.T) {
	inputFilename := "../test_files/image.gif"
	outputFilename := "../test_files/outputs/compressed_lossy.gif"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessGifOptions{OutputFilename: ""}
		err := functions.CompressGIFLossless(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("invalid quality option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyGifOptions{Quality: 0, OutputFilename: outputFilename}
		err := functions.CompressGIFLossy(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid quality option")
		}

		options = functions.LossyGifOptions{Quality: 101, OutputFilename: outputFilename}
		err = functions.CompressGIFLossy(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid quality option")
		}
	})

	t.Run("compression with no specified options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyGifOptions{}
		err := functions.CompressGIFLossy(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyGifOptions{OutputFilename: outputFilename, Quality: 80}
		err := functions.CompressGIFLossy(inputFilename, options)
		if err != nil {
			t.Error("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossy compression did not reduce the file size")
		}
	})
}
