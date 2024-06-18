package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// Options for SVG compression
type SvgOptions struct {
	OutputFilename string // Output filename
}

// CompressSVG compresses an SVG image using SVGO.
func CompressSVG(inputFilename string, options SvgOptions) error {
	// Verify input filename file exists
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Verify output file name
	if options.OutputFilename == "" {
		return fmt.Errorf("invalid output filename")
	}

	// Execute command
	cmd := exec.Command("svgo", inputFilename, "-o", options.OutputFilename)
	err := cmd.Run()
	return err
}
