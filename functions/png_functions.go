package functions

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Options for lossless PNG compression (OptiPNG)
type LosslessPngOptions struct {
	Level          int    // Optimization level (0-7)
	Strip          bool   // Strip metadata
	Interlace      bool   // Interlace the image
	OutputFilename string // Output filename
}

// Options for lossy PNG compression (Pngquant)
type LossyPngOptions struct {
	Quality        string // Quality range (e.g., "60-80")
	Speed          int    // Speed (1-11, higher is faster but less accurate)
	OutputFilename string // Output filename
}

// Function to perform lossless PNG compression
func CompressPNGLossless(inputFilename string, options LosslessPngOptions) error {
	// Verify output file name
	if options.OutputFilename == "" {
		return fmt.Errorf("invalid output filename")
	}

	// Create command arguments array
	args := []string{inputFilename}

	// Handle default optimization level
	if options.Level == 0 {
		options.Level = 1
	}

	// Handle optimization level
	if options.Level < 1 || options.Level > 7 {
		return fmt.Errorf("invalid optimization level (0-7)")
	}
	args = append(args, fmt.Sprintf("-o%d", options.Level))

	// Handle strip option
	if options.Strip {
		args = append(args, "-strip", "all")
	}

	// Handle interlace option
	if options.Interlace {
		args = append(args, "-i", "1")
	}

	// Handle output filename
	args = append(args, "-out", options.OutputFilename)

	cmd := exec.Command("optipng", args...) // Run OptiPNG
	err := cmd.Run()
	return err
}

// Function to perform lossy PNG compression
func CompressPNGLossy(inputFilename string, options LossyPngOptions) error {
	// Verify quality string is valid (e.g., "60-80")
	qualityParts := strings.Split(options.Quality, "-")
	if len(qualityParts) != 2 {
		return fmt.Errorf("invalid quality range")
	}
	qualityMin, err := strconv.Atoi(qualityParts[0])
	if err != nil {
		return fmt.Errorf("invalid quality range")
	}
	qualityMax, err := strconv.Atoi(qualityParts[1])
	if err != nil {
		return fmt.Errorf("invalid quality range")
	}
	if qualityMin < 0 || qualityMin > 100 || qualityMax < 0 || qualityMax > 100 || qualityMin > qualityMax {
		return fmt.Errorf("invalid quality range")
	}

	// Verify output filename
	if options.OutputFilename == "" {
		return fmt.Errorf("invalid output filename")
	}

	// Create command arguments array
	args := []string{"--quality", options.Quality}
	if options.Speed != 0 { // Default to 3 if not specified
		args = append(args, "--speed", strconv.Itoa(options.Speed))
	}
	args = append(args, inputFilename)

	// Handle output filename
	args = append(args, "--output", options.OutputFilename)

	cmd := exec.Command("pngquant", args...) // Run Pngquant
	err = cmd.Run()
	return err
}
