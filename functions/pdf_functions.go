package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// Options for PDF compression (Ghostscript)
type PdfOptions struct {
	Quality        int    // Quality level (0-100)
	OutputFilename string // Output filename
}

// Function to perform PDF compression
func CompressPDF(inputFilename string, options PdfOptions) error {
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
		"-sDEVICE=pdfwrite",
		"-dCompatibilityLevel=1.4",
		"-dPDFSETTINGS=/" + getPdfSettings(options.Quality),
		"-dNOPAUSE",
		"-dQUIET",
		"-dBATCH",
		"-sOutputFile=" + options.OutputFilename,
		inputFilename,
	}

	cmd := exec.Command("gs", args...)
	err := cmd.Run()
	return err
}

// Helper function to map quality to PDF settings
func getPdfSettings(quality int) string {
	switch {
	case quality < 25:
		return "screen"
	case quality < 50:
		return "ebook"
	case quality < 75:
		return "printer"
	default:
		return "prepress"
	}
}
