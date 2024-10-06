package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kingtingthegreat/ansi-converter/converter"
)

const DEFAULT_DIM = 40.
const DEFAULT_CHAR = "█"
const DEFAULT_RATIO = 0.46

const FILE = "file="
const DIM = "dim="
const CHAR = "char="
const RATIO = "ratio="

var OPTS = []string{FILE, DIM}

func parse_args() (string, float64, string, float64, error) {
	var filePath string
	dim := DEFAULT_DIM
	char := DEFAULT_CHAR
	ratio := DEFAULT_RATIO

	for _, arg := range os.Args[1:] {
		for _, opt := range OPTS {
			if strings.HasPrefix(arg, opt) {
				val := strings.TrimLeft(arg, opt)
				switch opt {
				case FILE:
					filePath = val
				case DIM:
					newDim, err := strconv.ParseFloat(val, 64)
					if err == nil {
						dim = newDim
					}
				case CHAR:
					char = string(val[0])
				case RATIO:
					newRatio, err := strconv.ParseFloat(val, 64)
					if err == nil {
						ratio = newRatio
					}
				}
			}
		}
	}

	return filePath, dim, char, ratio, nil
}

func main() {
	filePath, dim, char, ratio, err := parse_args()
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("error opening file")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("error decoding image")
	}

	x := converter.Convert(img, dim, char, ratio)
	fmt.Println(x)
}
