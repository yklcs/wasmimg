package cram_test

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	"github.com/yklcs/cram"
)

//go:embed example.jpeg
var raw []byte

func ExampleMozJPEG() {
	fmt.Printf("%d bytes raw\n", len(raw))

	img, err := jpeg.Decode(bytes.NewReader(raw))
	if err != nil {
		log.Fatalln(err)
	}

	b := img.Bounds()
	rgb := imageToRGB(img)

	compressed, err := cram.MozJPEG(rgb, b.Dx(), b.Dy())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%d bytes compressed\n", len(compressed))
	os.WriteFile("encoded.jpeg", compressed, 0644)

	// Output:
	// 2585457 bytes raw
	// 279145 bytes compressed
}

// imageToRGB converts img to a RGB byte slice.
func imageToRGB(img image.Image) []byte {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	rgb := make([]byte, width*height*3)
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgbaIndex := (y*width + x) * 4
			rgbIndex := (y*width + x) * 3
			pix := rgba.Pix[rgbaIndex : rgbaIndex+4]
			rgb[rgbIndex] = pix[0]
			rgb[1] = pix[1]
			rgb[2] = pix[2]
			copy(rgb[rgbIndex:rgbIndex+3], pix)
		}
	}

	return rgb
}
