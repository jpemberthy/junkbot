// Playing around with https://maxhalford.github.io/blog/halftoning-1
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var black = color.Gray{0}
var white = color.Gray{255}

func main() {
	img, _ := loadImage("img/junkbot.png")
	gray := rgbaToGray(img)
	dithered := thresholdDither(gray)

	f, _ := os.Create("dithered.png")
	defer f.Close()
	png.Encode(f, dithered)
}

func thresholdDither(gray *image.Gray) *image.Gray {
	bounds := gray.Bounds()
	dithered := image.NewGray(bounds)
	width := bounds.Dx()
	height := bounds.Dy()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			var c = blackOrWhite(gray.GrayAt(i, j))
			dithered.SetGray(i, j, c)
		}
	}
	return dithered
}

func loadImage(filepath string) (image.Image, error) {
	infile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer infile.Close()
	img, _, err := image.Decode(infile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func rgbaToGray(img image.Image) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}

	return gray
}

func blackOrWhite(g color.Gray) color.Gray {
	if g.Y < 123 {
		return black
	}
	return white
}
