package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// Options for lossless GIF compression
type LosslessGifOptions struct {
	OutputFilename string // Output filename
}

// Options for lossy GIF compression
type LossyGifOptions struct {
	Quality        int    // Quality level (1-100)
	OutputFilename string // Output filename
}

// CompressGIFLossless compresses a GIF losslessly using Gifsicle.
// Parameters:
//   - inputFilename:  The path to the input GIF file.
//   - options: Gifsicle-specific options (e.g., "--lossy=80", "--optimize").
func CompressGIFLossless(inputFilename string, options LosslessGifOptions) error {
	// Verify input filename file exists
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Verify output file name
	if options.OutputFilename == "" {
		return fmt.Errorf("invalid output filename")
	}

	// Execute command
	cmd := exec.Command("gifsicle", "--optimize", inputFilename, "-o", options.OutputFilename)
	err := cmd.Run()
	return err
}

// CompressGIFLossless compresses a GIF lossy using Gifsicle.
// Parameters:
//   - inputFilename:  The path to the input GIF file.
//   - options: Gifsicle-specific options (e.g., "--lossy=80", "--optimize").
func CompressGIFLossy(inputFilename string, options LossyGifOptions) error {
	// Verify input filename file exists
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Verify output file name
	if options.OutputFilename == "" {
		return fmt.Errorf("invalid output filename")
	}

	// Verify quality level
	if options.Quality < 1 || options.Quality > 100 {
		return fmt.Errorf("invalid quality level (1-100)")
	}

	// Execute command
	cmd := exec.Command("gifsicle", fmt.Sprintf("--lossy=%d", options.Quality), inputFilename, "-o", options.OutputFilename)
	fmt.Println(cmd.String())
	err := cmd.Run()
	return err
}
