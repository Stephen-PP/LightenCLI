package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressTIFFLossless(t *testing.T) {
	inputFilename := "../test_files/image.tiff"
	outputFilename := "../test_files/outputs/compressed_lossless.tiff"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid input file", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessTiffOptions{OutputFilename: outputFilename, Compression: "LZW", Depth: 8}
		err := functions.CompressTIFFLossless("non-existent-file.tiff", options)
		if err == nil {
			t.Error("Expected error for non-existent input file")
		}
	})

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessTiffOptions{OutputFilename: "", Compression: "LZW", Depth: 8}
		err := functions.CompressTIFFLossless(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("invalid compression option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessTiffOptions{OutputFilename: outputFilename, Compression: "invalid", Depth: 8}
		err := functions.CompressTIFFLossless(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid compression option")
		}
	})

	t.Run("invalid depth option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessTiffOptions{OutputFilename: outputFilename, Compression: "LZW", Depth: 0}
		err := functions.CompressTIFFLossless(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid depth option")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessTiffOptions{OutputFilename: outputFilename, Compression: "LZW", Depth: 8}
		err := functions.CompressTIFFLossless(inputFilename, options)
		if err != nil {
			t.Error("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossless TIFF compression did not reduce the file size")
		}
	})
}

func TestCompressTIFFLossy(t *testing.T) {
	inputFilename := "../test_files/image.tiff"
	outputFilename := "../test_files/outputs/compressed_lossy.tiff"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid input file", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyTiffOptions{OutputFilename: outputFilename, Quality: 80}
		err := functions.CompressTIFFLossy("non-existent-file.tiff", options)
		if err == nil {
			t.Error("Expected error for non-existent input file")
		}
	})

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyTiffOptions{OutputFilename: "", Quality: 80}
		err := functions.CompressTIFFLossy(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("invalid quality option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyTiffOptions{OutputFilename: outputFilename, Quality: -1}
		err := functions.CompressTIFFLossy(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid quality option")
		}

		options = functions.LossyTiffOptions{OutputFilename: outputFilename, Quality: 101}
		err = functions.CompressTIFFLossy(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid quality option")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyTiffOptions{OutputFilename: outputFilename, Quality: 80}
		err := functions.CompressTIFFLossy(inputFilename, options)
		if err != nil {
			t.Error("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossy TIFF compression did not reduce the file size")
		}
	})
}
