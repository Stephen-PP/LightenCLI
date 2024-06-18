package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Options for WebP compression
type WebPOptions struct {
	Quality        int    // Quality level (0-100)
	OutputFilename string // Output filename
}

// Function to perform WebP compression
func CompressWebPLossy(inputFilename string, options WebPOptions) error {
	// Verify input filename file exists
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Verify quality is valid (0-100)
	if options.Quality < 0 || options.Quality > 100 {
		return fmt.Errorf("invalid quality level (0-100)")
	}

	// Verify output file name
	if options.OutputFilename == "" {
		return fmt.Errorf("invalid output filename")
	}

	// Create command arguments array
	args := []string{
		"-q", strconv.Itoa(options.Quality),
		"-o", options.OutputFilename,
		inputFilename,
	}

	cmd := exec.Command("cwebp", args...)
	err := cmd.Run()
	return err
}

// Function to perform WebP lossless compression
func CompressWebPLossless(inputFilename string, options WebPOptions) error {
	// Verify input filename file exists
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Verify quality is valid (0-100)
	if options.Quality < 0 || options.Quality > 100 {
		return fmt.Errorf("invalid quality level (0-100)")
	}

	// Verify output file name
	if options.OutputFilename == "" {
		return fmt.Errorf("invalid output filename")
	}

	// Create command arguments array
	args := []string{
		"-lossless",
		"-q", strconv.Itoa(options.Quality),
		"-o", options.OutputFilename,
		inputFilename,
	}

	cmd := exec.Command("cwebp", args...)
	err := cmd.Run()
	return err
}
