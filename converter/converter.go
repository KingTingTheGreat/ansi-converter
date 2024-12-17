package converter

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strconv"
	"strings"

	"github.com/kingtingthegreat/ansi-converter/config"
	"github.com/nfnt/resize"
)

const ANSI_FOREGROUND = "\x1b[38;2;"
const ANSI_RESET = "\x1b[0m"

func RGBtoAnsi(r, g, b int) string {
	return ANSI_FOREGROUND + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}

func Convert(img image.Image, cfg *config.Config) string {
	// assume square image
	h := cfg.Dim * cfg.FontRatio
	img = resize.Resize(uint(cfg.Dim), uint(h), img, resize.Lanczos3)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	if width != int(cfg.Dim) {
		log.Fatal("does not match", cfg.Dim, width)
	}

	output := strings.Repeat("\n", cfg.PaddingTop)
	for y := 0; y < height; y++ {
		output += strings.Repeat(" ", cfg.PaddingLeft)
		for x := 0; x < width; x++ {
			pixel := img.At(x, y)
			r, g, b, _ := pixel.RGBA()
			r, g, b = r>>8, g>>8, b>>8 // convert to 8 bit color

			colorStr := RGBtoAnsi(int(r), int(g), int(b))

			output += colorStr + cfg.Char
		}
		output += strings.Repeat(" ", cfg.PaddingRight)
		if y == height-1 {
			output += ANSI_RESET
		} else {
			output += "\n"
		}
	}
	output += strings.Repeat("\n", cfg.PaddingBottom)

	return output
}
