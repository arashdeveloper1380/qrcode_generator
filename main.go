package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/skip2/go-qrcode"
)

func generateSVG(content string, size int) string {
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		log.Fatalf("Failed to create QR code: %v", err)
	}

	// QR code matrix
	matrix := qr.Bitmap()
	scale := size / len(matrix) // Scale each cell based on the matrix size
	var svg strings.Builder

	// SVG header
	svg.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`, size, size, size, size))

	// Generate SVG rectangles for each black cell
	for y, row := range matrix {
		for x, cell := range row {
			if cell {
				svg.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%d" height="%d" fill="black"/>`, x*scale, y*scale, scale, scale))
			}
		}
	}

	// SVG footer
	svg.WriteString("</svg>")
	return svg.String()
}

func main() {
	// Directory to save the QR codes
	outputDir := "qrcodes"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Loop to generate QR codes
	for i := 1; i <= 1000; i++ {
		// QR code content
		content := fmt.Sprintf("%d", i)

		// File name for the QR code
		filename := filepath.Join(outputDir, fmt.Sprintf("%d.svg", i))

		// Generate the SVG content
		svgContent := generateSVG(content, 48)

		// Save the SVG to a file
		err := os.WriteFile(filename, []byte(svgContent), 0644)
		if err != nil {
			log.Printf("Failed to save file %s: %v", filename, err)
			continue
		}

		fmt.Printf("Generated QR code: %s\n", filename)
	}

	fmt.Println("QR code generation complete!")
}
