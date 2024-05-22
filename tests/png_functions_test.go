package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressPNGLossless(t *testing.T) {
	inputFilename := "../test_files/image.png"
	outputFilename := "../test_files/outputs/compressed_lossless.png"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessPngOptions{Level: 6, Strip: true, OutputFilename: ""}
		err := functions.CompressPNGLossless(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("invalid optimization level", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessPngOptions{Level: 8, Strip: true, OutputFilename: outputFilename}
		err := functions.CompressPNGLossless(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid optimization level")
		}
	})

	t.Run("compression with default options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessPngOptions{OutputFilename: outputFilename}
		err := functions.CompressPNGLossless(inputFilename, options)
		if err != nil {
			t.Error("Lossless compression failed:", err)
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossless compression did not reduce the file size")
		}
	})

	t.Run("compression with custom options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LosslessPngOptions{Level: 6, Strip: true, OutputFilename: outputFilename}

		err := functions.CompressPNGLossless(inputFilename, options)
		if err != nil {
			t.Fatal("Lossless compression failed:", err)
		}

		// Compare file sizes
		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossless compression did not reduce file size")
		}
	})
}

func TestCompressPNGLossy(t *testing.T) {
	inputFilename := "../test_files/image.png"
	outputFilename := "../test_files/outputs/compressed_lossy.png"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyPngOptions{Quality: "65-80", Speed: 5, OutputFilename: ""}
		err := functions.CompressPNGLossy(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("compression with invalid quality range (lower bound)", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyPngOptions{Quality: "-2-80", Speed: 5, OutputFilename: outputFilename}
		err := functions.CompressPNGLossy(inputFilename, options)
		if err != nil && err.Error() != "invalid quality range" {
			t.Error("Expected error for invalid quality range")
		}
	})

	t.Run("compression with invalid quality range (upper bound)", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyPngOptions{Quality: "10-110", Speed: 5, OutputFilename: outputFilename}
		err := functions.CompressPNGLossy(inputFilename, options)
		if err != nil && err.Error() != "invalid quality range" {
			t.Error("Expected error for invalid quality range")
		}
	})

	t.Run("compression with invalid quality range (min > max)", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyPngOptions{Quality: "90-80", Speed: 5, OutputFilename: outputFilename}
		err := functions.CompressPNGLossy(inputFilename, options)
		if err != nil && err.Error() != "invalid quality range" {
			t.Error("Expected error for invalid quality range")
		}
	})

	t.Run("compression with default options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyPngOptions{Quality: "65-80", Speed: 5, OutputFilename: outputFilename}
		err := functions.CompressPNGLossy(inputFilename, options)
		if err != nil {
			t.Error("Lossy compression failed:", err)
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossy compression did not reduce file size")
		}
	})

	t.Run("compression with custom options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.LossyPngOptions{Quality: "65-80", Speed: 5, OutputFilename: outputFilename}
		err := functions.CompressPNGLossy(inputFilename, options)
		if err != nil {
			t.Error("Lossy compression failed:", err)
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("Lossy compression did not reduce file size")
		}
	})
}
