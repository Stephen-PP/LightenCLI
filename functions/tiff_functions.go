package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// LosslessTiffOptions provides options for lossless TIFF compression.
type LosslessTiffOptions struct {
	OutputFilename string // Output filename
	Compression    string // Compression type (e.g., "LZW", "ZIP", "NONE")
	Depth          int    // Bit depth (1, 8, 16, 32)
}

// CompressTIFFLossless compresses a TIFF image losslessly using ImageMagick.
func CompressTIFFLossless(inputFilename string, options LosslessTiffOptions) error {
	// Input file validation
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Validate input options
	if options.OutputFilename == "" || (options.Compression != "LZW" && options.Compression != "ZIP" && options.Compression != "NONE") || (options.Depth != 1 && options.Depth != 8 && options.Depth != 16 && options.Depth != 32) {
		return fmt.Errorf("invalid options")
	}

	// Build the ImageMagick command with options
	args := []string{
		inputFilename,
		"-compress", options.Compression,
		"-depth", strconv.Itoa(options.Depth),
		options.OutputFilename,
	}

	// Execute ImageMagick
	cmd := exec.Command("convert", args...)
	err := cmd.Run()
	return err
}

// LossyTiffOptions provides options for lossy TIFF compression.
type LossyTiffOptions struct {
	OutputFilename string // Output filename
	Quality        int    // JPEG quality level (0-100)
}

// CompressTIFFLossy compresses a TIFF image lossily using ImageMagick's JPEG compression.
func CompressTIFFLossy(inputFilename string, options LossyTiffOptions) error {
	// Input file validation
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Validate input options
	if options.OutputFilename == "" || (options.Quality < 0 || options.Quality > 100) {
		return fmt.Errorf("invalid options")
	}

	// Build the ImageMagick command
	args := []string{
		inputFilename,
		"-quality", fmt.Sprintf("%d", options.Quality),
		options.OutputFilename,
	}

	// Execute ImageMagick
	cmd := exec.Command("convert", args...)
	err := cmd.Run()
	return err
}
