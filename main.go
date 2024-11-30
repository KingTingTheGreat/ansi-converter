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

const FILE = "file"
const DIM = "dim"
const CHAR = "char"
const RATIO = "ratio"

type Config struct {
	FilePath  string
	Dim       float64
	Char      string
	FontRatio float64
}

func parse_args(cfg Config) (Config, error) {
	for _, argValStr := range os.Args[1:] {
		argVal := strings.SplitN(argValStr, "=", 2)
		arg := argVal[0]
		if len(argVal) == 1 {
			cfg.FilePath = arg
		} else {
			val := argVal[1]
			switch arg {
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
