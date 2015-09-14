package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

func main() {

	// Maximum width and height of the displayed image, measured in chars
	maxWidth := 80
	maxHeight := 80

	// The ratio of height to width for a single character, this also influences maxHeight.
	// 0.5 means that 1 width == 0.5 height
	ratio := 0.5

	if len(os.Args) < 2 {
		fmt.Printf("Please give filename.\ni.e.\nimgview image.ext\n")
		os.Exit(1)
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	// TODO: do the resize in one single call
	ratioImg := resize.Resize(uint(img.Bounds().Dx()), uint(float64(img.Bounds().Dy())*ratio), img, resize.Lanczos3)
	thumb := resize.Thumbnail(uint(maxWidth), uint(maxHeight), ratioImg, resize.Lanczos3)

	imgWidth := thumb.Bounds().Dx()
	imgHeight := thumb.Bounds().Dy()
	greys := strings.Split(" .:-=+*#@", "")
	numGreys := len(greys)
	divideBy := 3 * (65536 / numGreys)
	outStr := ""
	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			r, g, b, _ := thumb.At(x, y).RGBA()
			result := int(r+g+b) / divideBy
			if result >= numGreys {
				result = numGreys - 1
			}
			outStr += greys[result]
		}
		outStr += "\n"
	}
	fmt.Println(outStr)
}
