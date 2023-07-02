package wasmimg_test

import (
	"bytes"
	_ "embed"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"testing"

	"github.com/yklcs/wasmimg/mozjpeg"
)

//go:embed example.jpeg
var raw []byte

func TestMozjpeg(t *testing.T) {
	t.Logf("%d bytes raw\n", len(raw))

	img, err := jpeg.Decode(bytes.NewReader(raw))
	if err != nil {
		log.Fatalln(err)
	}
	b := img.Bounds()

	decompressed, err := mozjpeg.Decode(raw)
	if err != nil {
		log.Fatalln(err)
	}
	t.Logf("%d bytes decompressed\n", len(decompressed))

	compressed, err := mozjpeg.Encode(decompressed, b.Dx(), b.Dy(), 75)
	if err != nil {
		log.Fatalln(err)
	}

	var buf bytes.Buffer
	jpeg.Encode(&buf, img, &jpeg.Options{Quality: 75})
	t.Logf("%d bytes compressed via image/jpeg\n", buf.Len())

	t.Logf("%d bytes compressed\n", len(compressed))
	os.WriteFile("encoded.jpeg", compressed, 0755)
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
