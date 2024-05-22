package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Options for JPEG compression (MozJPEG)
type JpegOptions struct {
	Quality        int    // Quality level (1-100)
	OutputFilename string // Output filename
}

// Function to perform JPEG compression
func CompressJPEG(inputFilename string, options JpegOptions) error {
	// Verify input filename file exists
	if _, err := os.Stat(inputFilename); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist")
	}

	// Verify quality is valid (1-100)
	if options.Quality < 1 || options.Quality > 100 {
		return fmt.Errorf("invalid quality level (1-100)")
	}

	// Verify output file name
	if options.OutputFilename == "" {
		return fmt.Errorf("invalid output filename")
	}

	cmdOne := exec.Command("djpeg", inputFilename)
	cmdTwo := exec.Command("cjpeg", "-quality", strconv.Itoa(options.Quality), "-optimize", "-outfile", options.OutputFilename)

	cmdTwo.Stdin, _ = cmdOne.StdoutPipe()

	_ = cmdTwo.Start()
	_ = cmdOne.Run()
	err := cmdTwo.Wait()

	return err
}
