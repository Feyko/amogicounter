package main

import (
	"analysus/amogi"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func newAmogus(width, height int) amogi.Amogus {
	var amogus amogi.Amogus
	for i := 0; i < height; i++ {
		amogus = append(amogus, make([]amogi.Pixel, width))
	}
	return amogus
}

func amogusImgToAmogusTemplate(img image.Image) amogi.Amogus {
	maxY := img.Bounds().Max.Y
	maxX := img.Bounds().Max.X
	amogus := newAmogus(maxX, maxY)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			amogus[y][x] = img.At(x, y) == color.NRGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			}
		}
	}
	return amogus
}

func cleanOutput(output string) string {
	data := []byte(output)
	noVerbose := regexp.MustCompile(`(\[])?amogi\.(Amogus|Pixel)`)
	data = noVerbose.ReplaceAll(data, []byte{})
	noFlat := regexp.MustCompile(`, {`)
	data = noFlat.ReplaceAll(data, []byte(", \n{"))
	spaceStart := regexp.MustCompile(`{{`)
	data = spaceStart.ReplaceAll(data, []byte("{\n{"))
	spaceEnd := regexp.MustCompile(`}}`)
	data = spaceEnd.ReplaceAll(data, []byte("},\n}"))
	revalidate := regexp.MustCompile(`}\n`)
	data = revalidate.ReplaceAll(data, []byte("},\n"))
	return string(data)
}

func main() {
	inputPathPtr := flag.String("i", "", "Path to the directory containing the templates")
	outputPathPtr := flag.String("o", "", "Where to write the output")
	flag.Parse()
	inputPath := *inputPathPtr
	outputPath := *outputPathPtr
	if inputPath == "" {
		log.Fatalln("Argument 'i' required")
	}
	toConsole := false
	if outputPath == "" {
		toConsole = true
	}
	output := ""
	err := filepath.Walk(inputPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		imageReader, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("could not open the input image at %v: %w\n", path, err)
		}
		img, err := png.Decode(imageReader)
		if err != nil {
			return fmt.Errorf("could not decode the image at %v. Make sure it is a PNG: %w", path, err)
		}
		template := amogusImgToAmogusTemplate(img)
		output += fmt.Sprintf("%#v\n", template)
		return nil
	})
	output = cleanOutput(output)
	if err != nil {
		log.Fatal(err)
	}
	if toConsole {
		fmt.Println(output)
		os.Exit(0)
	}
	err = os.WriteFile(outputPath, []byte(output), 0644)
}
