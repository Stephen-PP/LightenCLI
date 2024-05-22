package main

import (
	"fmt"

	functions "github.com/Stephen-PP/LightenCLI/functions"
)

func main() {
	// Lossless
	losslessOpts := functions.LosslessPngOptions{Level: 6, Strip: true, OutputFilename: "./test_files/outputs/image-lossless.png"}
	err := functions.CompressPNGLossless("./test_files/image.png", losslessOpts)
	if err != nil {
		fmt.Println(err)
	}

	// Lossy
	lossyOpts := functions.LossyPngOptions{Quality: "65-80", Speed: 5, OutputFilename: "./test_files/outputs/image-lossy.png"}
	err = functions.CompressPNGLossy("./test_files/image.png", lossyOpts)
	if err != nil {
		fmt.Println(err)
	}

	// JPEG
	jpegOpts := functions.JpegOptions{Quality: 80, OutputFilename: "./test_files/outputs/compressed_jpeg.jpg"}
	err = functions.CompressJPEG("./test_files/image.jpg", jpegOpts)
	if err != nil {
		fmt.Println(err)
	}

	// PDF
	pdfOpts := functions.PdfOptions{Quality: 75, OutputFilename: "./test_files/outputs/compressed_pdf.pdf"}
	err = functions.CompressPDF("./test_files/document.pdf", pdfOpts)
	if err != nil {
		fmt.Println(err)
	}
}
