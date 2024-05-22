package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressPDF(t *testing.T) {
	inputFilename := "../test_files/document.pdf"
	outputFilename := "../test_files/outputs/compressed_pdf.pdf"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.PdfOptions{Quality: 75, OutputFilename: ""}
		err := functions.CompressPDF(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid output filename")
		}
	})

	t.Run("invalid quality level", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.PdfOptions{Quality: -1, OutputFilename: outputFilename}
		err := functions.CompressPDF(inputFilename, options)
		if err == nil {
			t.Error("Expected error for invalid quality level")
		}
	})

	t.Run("compression with default options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.PdfOptions{Quality: 75, OutputFilename: outputFilename}
		err := functions.CompressPDF(inputFilename, options)
		if err != nil {
			t.Error("PDF compression failed:", err)
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() >= inputFileInfo.Size() {
			t.Error("PDF compression did not reduce the file size")
		}
	})
}
