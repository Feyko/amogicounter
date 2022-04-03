package main

import (
	"analysus/amogi"
	"analysus/highlight"
	"flag"
	"fmt"
	"github.com/disintegration/imaging"
	_ "image/png"
	"log"
)

func main() {
	inputPathPtr := flag.String("i", "", "Path to the image to scan")
	outputPathPtr := flag.String("o", "", "Where to write the output")
	flag.Parse()
	inputPath := *inputPathPtr
	outputPath := *outputPathPtr
	if inputPath == "" {
		log.Fatalln("Argument 'i' required")
	}
	if outputPath == "" {
		log.Fatalln("Argument 'o' required")
	}
	img, err := imaging.Open(inputPath)
	if err != nil {
		log.Fatalf("Could not open the input image: %w", err)
	}
	foundAmogi := amogi.ScanAmogi(img)
	fmt.Println(len(foundAmogi))
	highlighted := highlight.Highlight(img, foundAmogi, 0.7)
	err = imaging.Save(highlighted, outputPath)
	if err != nil {
		log.Fatalf("Could not write the output image: %w", err)
	}
}
