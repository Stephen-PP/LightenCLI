package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// MP3Options encapsulates settings for the "lame" MP3 encoder.
type MP3Options struct {
	Bitrate        int    // Target bitrate in kbps (e.g., 128, 192, 320)
	Quality        int    // Quality setting (0=best, 9=worst) – optional
	DesiredBitrate int    // Desired bitrate in kbit/s – optional (override quality)
	OutputFilename string // Name of the output MP3 file
}

// CompressMP3 compresses an input MP3 file using the "lame" encoder.
func CompressMP3(inputFilename string, options MP3Options) error {
	// Input File Existence Check
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file '%s' does not exist", inputFilename)
	}

	// Output File Validation
	if options.OutputFilename == "" {
		// Default output filename if not provided (e.g., add ".mp3" to input)
		options.OutputFilename = inputFilename + ".mp3"
	}

	// Bitrate Validation
	if options.Bitrate != 8 && options.Bitrate != 16 && options.Bitrate != 24 &&
		options.Bitrate != 32 && options.Bitrate != 40 && options.Bitrate != 48 &&
		options.Bitrate != 56 && options.Bitrate != 64 && options.Bitrate != 80 &&
		options.Bitrate != 96 && options.Bitrate != 112 && options.Bitrate != 128 &&
		options.Bitrate != 160 && options.Bitrate != 192 && options.Bitrate != 224 &&
		options.Bitrate != 256 && options.Bitrate != 320 {
		return fmt.Errorf("invalid bitrate (valid options: 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320)")
	}

	// Base Command Construction
	cmdArgs := []string{
		"-b", fmt.Sprintf("%d", options.Bitrate), // Set bitrate
		inputFilename, options.OutputFilename, // Input and output files
	}

	// Optional Settings
	if options.DesiredBitrate != 0 { // If desired file size is specified
		// Calculate average bitrate required to achieve desired file size
		// Formula: Bitrate = (DesiredFileSize * 8) / DurationInSeconds
		cmdArgs = append(cmdArgs, "--abr", fmt.Sprintf("%d", options.DesiredBitrate))
	} else if options.Quality != 0 && options.Quality <= 9 { // If quality is specified
		cmdArgs = append(cmdArgs, "-V", fmt.Sprintf("%d", options.Quality))
	}

	// Compression Execution
	cmd := exec.Command("lame", cmdArgs...) // Create the "lame" command
	output, err := cmd.CombinedOutput()     // Run it and capture output

	// Error Handling
	if err != nil {
		return fmt.Errorf("MP3 compression failed: %s (%s)", err, string(output))
	}

	return nil
}
