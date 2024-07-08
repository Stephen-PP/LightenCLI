package tests

import (
	"os"
	"testing"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func TestCompressMP3(t *testing.T) {
	inputFilename := "../test_files/sample.mp3"
	outputFilename := "../test_files/outputs/compressed.mp3"

	// Clean up before testing
	os.Remove(outputFilename)

	t.Run("invalid input file", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.MP3Options{Bitrate: 128, OutputFilename: outputFilename}
		err := functions.CompressMP3("non-existent-file.mp3", options)
		if err == nil {
			t.Error("Expected error for non-existent input file")
		}
	})

	t.Run("invalid output filename", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.MP3Options{Bitrate: 128, OutputFilename: ""}
		err := functions.CompressMP3(inputFilename, options)
		if err == nil {
			t.Errorf("Expected error for default output filename, got nil")
		}
	})

	t.Run("invalid bitrate option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.MP3Options{Bitrate: 0, OutputFilename: outputFilename}
		err := functions.CompressMP3(inputFilename, options)

		if err == nil {
			t.Fatalf("Expected error for invalid bitrate option")
		}

		options = functions.MP3Options{Bitrate: 512, OutputFilename: outputFilename}
		err = functions.CompressMP3(inputFilename, options)

		if err == nil {
			t.Error("Expected error for invalid bitrate option")
		}
	})

	t.Run("compression with valid options", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.MP3Options{Bitrate: 128, OutputFilename: outputFilename}
		err := functions.CompressMP3(inputFilename, options)
		if err != nil {
			t.Fatalf("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() > inputFileInfo.Size() {
			t.Error("MP3 compression did not reduce the file size")
		}
	})

	t.Run("compression with desired bitrate", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.MP3Options{DesiredBitrate: 1000, OutputFilename: outputFilename}
		err := functions.CompressMP3(inputFilename, options)
		if err != nil {
			t.Fatalf("Expected no error for valid options: %v", err)
			return
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() > inputFileInfo.Size() {
			t.Error("MP3 compression did not reduce the file size")
		}
	})

	t.Run("compression with quality option", func(t *testing.T) {
		defer os.Remove(outputFilename) // After test clean up

		options := functions.MP3Options{Quality: 5, OutputFilename: outputFilename}
		err := functions.CompressMP3(inputFilename, options)
		if err != nil {
			t.Fatalf("Expected no error for valid options")
		}

		inputFileInfo, _ := os.Stat(inputFilename)
		outputFileInfo, _ := os.Stat(outputFilename)

		if outputFileInfo.Size() > inputFileInfo.Size() {
			t.Error("MP3 compression did not reduce the file size")
		}
	})
}
