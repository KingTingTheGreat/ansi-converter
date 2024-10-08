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
	"github.com/kingtingthegreat/ansi-converter/defaults"
)

const FILE = "file="
const DIM = "dim="
const CHAR = "char="
const RATIO = "ratio="

var OPTS = []string{FILE, DIM, CHAR, RATIO}

type Config struct {
	FilePath  string
	Dim       float64
	Char      string
	FontRatio float64
}

func parse_args(cfg Config) (Config, error) {
	for i, arg := range os.Args[1:] {
		isOpt := false
		for _, opt := range OPTS {
			if strings.HasPrefix(arg, opt) {
				isOpt = true
				val := strings.TrimLeft(arg, opt)
				switch opt {
				case FILE:
					cfg.FilePath = val
				case DIM:
					newDim, err := strconv.ParseFloat(val, 64)
					if err == nil {
						cfg.Dim = newDim
					} else {
						return Config{}, fmt.Errorf("invalid dim")
					}
				case CHAR:
					if len(val) == 0 {
						return Config{}, fmt.Errorf("invalid char")
					}
					cfg.Char = string(val[0])
				case RATIO:
					newRatio, err := strconv.ParseFloat(val, 64)
					if err == nil {
						cfg.FontRatio = newRatio
					} else {
						return Config{}, fmt.Errorf("invalid ratio")
					}
				}
			}
		}
		if i == 0 && !isOpt {
			cfg.FilePath = arg
		}
	}

	return cfg, nil
}

func main() {
	config := Config{
		FilePath:  "",
		Dim:       defaults.DEFAULT_DIM,
		Char:      defaults.DEFAULT_CHAR,
		FontRatio: defaults.DEFAULT_RATIO,
	}
	config, err := parse_args(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := os.Open(config.FilePath)
	if err != nil {
		log.Fatal("error opening file")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("error decoding image")
	}

	x := converter.Convert(img, config.Dim, config.Char, config.FontRatio)
	fmt.Println(x)
}
