package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// BMP compression is inherently lossless, so we only need one function.

// BMPOptions contains options for BMP compression.
type BMPOptions struct {
	OutputFilename string // Output filename
	Depth          int    // Bit depth (1, 4, 8, 16, 24, 32)
}

// CompressBMP compresses a BMP image using ImageMagick.
func CompressBMP(inputFilename string, options BMPOptions) error {
	// Input file validation
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Validate input options
	if options.OutputFilename == "" || (options.Depth != 1 && options.Depth != 4 && options.Depth != 8 && options.Depth != 16 && options.Depth != 24 && options.Depth != 32) {
		return fmt.Errorf("invalid output filename")
	}

	args := []string{
		inputFilename,
		options.OutputFilename,
	}

	// Execute ImageMagick
	cmd := exec.Command("convert", args...)
	err := cmd.Run()

	return err
}
