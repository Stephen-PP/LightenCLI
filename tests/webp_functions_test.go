package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressWebPLossy(t *testing.T) {
	inputFilename := "../test_files/image.webp"
	outputFilename := "../test_files/outputs/compressed_lossy.webp"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid input file", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.WebPOptions{Quality: 80, OutputFilename: outputFilename}
		err := functions.CompressWebPLossy("non-existent-file.webp", options)
		if err == nil {
			t.Error("Expected error for non-existent input file")
		}
	})

	t.Run("invalid quality option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.WebPOptions{Quality: -1, OutputFilename: outputFilename}
		err := functions.CompressWebPLossy(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid quality option")
		}

		options = functions.WebPOptions{Quality: 101, OutputFilename: outputFilename}
		err = functions.CompressWebPLossy(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid quality option")
		}
	})

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.WebPOptions{Quality: 80, OutputFilename: ""}
		err := functions.CompressWebPLossy(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.WebPOptions{Quality: 80, OutputFilename: outputFilename}
		err := functions.CompressWebPLossy(inputFilename, options)
		if err != nil {
			t.Error("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossy WebP compression did not reduce the file size")
		}
	})
}

func TestCompressWebPLossless(t *testing.T) {
	inputFilename := "../test_files/image.webp"
	outputFilename := "../test_files/outputs/compressed_lossless.webp"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid input file", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.WebPOptions{Quality: 80, OutputFilename: outputFilename}
		err := functions.CompressWebPLossless("non-existent-file.webp", options)
		if err == nil {
			t.Error("Expected error for non-existent input file")
		}
	})

	t.Run("invalid quality option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.WebPOptions{Quality: -1, OutputFilename: outputFilename}
		err := functions.CompressWebPLossless(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid quality option")
		}

		options = functions.WebPOptions{Quality: 101, OutputFilename: outputFilename}
		err = functions.CompressWebPLossless(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid quality option")
		}
	})

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.WebPOptions{Quality: 80, OutputFilename: ""}
		err := functions.CompressWebPLossless(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.WebPOptions{Quality: 30, OutputFilename: outputFilename}
		err := functions.CompressWebPLossless(inputFilename, options)
		if err != nil {
			t.Error("Expected no error for valid options")
		}

		if _, err := os.Stat(outputFilename); os.IsNotExist(err) {
			t.Error("Lossless WebP compression did not create the output file")
		}
	})
}
