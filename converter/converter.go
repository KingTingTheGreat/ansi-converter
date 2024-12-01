package converter

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

const ANSI_FOREGROUND = "\x1b[38;2;"
const ANSI_RESET = "\x1b[0m"

func RGBtoAnsi(r, g, b int) string {
	return ANSI_FOREGROUND + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}

func Convert(img image.Image, char string, dim, fontRatio float64, paddingTop, paddingRight, paddingBottom, paddingLeft int) string {
	// assume square image
	h := dim * fontRatio
	img = resize.Resize(uint(dim), uint(h), img, resize.Lanczos3)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	output := strings.Repeat("\n", paddingTop)
	for y := 0; y < height; y++ {
		output += strings.Repeat(" ", paddingLeft)
		for x := 0; x < width; x++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			r, g, b = r>>8, g>>8, b>>8 // convert to 8 bit color

			colorStr := RGBtoAnsi(int(r), int(g), int(b))

			output += colorStr + char
		}
		output += strings.Repeat(" ", paddingRight)
		if y != height-1 {
			output += ANSI_RESET + "\n"
		}
	}
	output += strings.Repeat("\n", paddingBottom)

	return output
}
