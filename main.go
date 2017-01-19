// Playing around with https://maxhalford.github.io/blog/halftoning-1
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

var black = color.Gray{0}
var white = color.Gray{255}

func main() {
	img, _ := loadImage("img/junkbot.png")
	gray := rgbaToGray(img)
	// dithered := thresholdDither(gray)
	dithered := gridDither(gray, 10, 8)

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

func gridDither(gray *image.Gray, k int, gamma float64) *image.Gray {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	bounds := gray.Bounds()
	dithered := newWhite(bounds)
	width := bounds.Dx()
	height := bounds.Dy()

	for i := 0; i < width; i += k {
		for j := 0; j < height; j += k {
			cell := rgbaToGray(gray.SubImage(image.Rect(i, j, i+k, j+k)))
			mu := avgIntensity(cell)
			n := (1 - mu) * gamma

			// TODO: double check this.
			for k := 1; k <= int(n); k++ {
				x := randInt(i, min(i+k, width), rng)
				y := randInt(j, min(j+k, height), rng)

				dithered.SetGray(x, y, black)
			}
		}
	}

	return dithered
}

func randInt(min, max int, rng *rand.Rand) int {
	return rng.Intn(max-min) + min
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

func newWhite(bounds image.Rectangle) *image.Gray {
	white := image.NewGray(bounds)
	for i := range white.Pix {
		white.Pix[i] = 255
	}
	return white
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

func avgIntensity(gray *image.Gray) float64 {
	var sum float64
	for _, pix := range gray.Pix {
		sum += float64(pix)
	}
	return sum / float64(len(gray.Pix)*256)
}

func blackOrWhite(g color.Gray) color.Gray {
	if g.Y < 123 {
		return black
	}
	return white
}
