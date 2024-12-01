package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kingtingthegreat/ansi-converter/defaults"
)

const FILE = "file"
const DIM = "dim"
const CHAR = "char"
const RATIO = "ratio"
const PADDING = "p"
const PADDING_TOP = "pT"
const PADDING_RIGHT = "pR"
const PADDING_BOTTOM = "pB"
const PADDING_LEFT = "pL"

type Config struct {
	FilePath      string
	Dim           float64
	Char          string
	FontRatio     float64
	PaddingTop    int
	PaddingRight  int
	PaddingBottom int
	PaddingLeft   int
}

var cfg = Config{
	FilePath:  "",
	Dim:       defaults.DEFAULT_DIM,
	Char:      defaults.DEFAULT_CHAR,
	FontRatio: defaults.DEFAULT_RATIO,
}

func ParseArgs() error {
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
					return fmt.Errorf("invalid dim")
				}
			case CHAR:
				if len(val) == 0 {
					return fmt.Errorf("invalid char")
				}
				cfg.Char = string(val[0])
			case RATIO:
				newRatio, err := strconv.ParseFloat(val, 64)
				if err == nil {
					cfg.FontRatio = newRatio
				} else {
					return fmt.Errorf("invalid ratio")
				}
			case PADDING:
				newPadding, err := strconv.Atoi(val)
				if err != nil {
					return fmt.Errorf("invalid padding")
				}
				cfg.PaddingTop = newPadding
				cfg.PaddingRight = newPadding
				cfg.PaddingBottom = newPadding
				cfg.PaddingLeft = newPadding
			case PADDING_TOP:
				newPaddingTop, err := strconv.Atoi(val)
				if err != nil {
					return fmt.Errorf("invalid padding top")
				}
				cfg.PaddingTop = newPaddingTop
			case PADDING_RIGHT:
				newPaddingRight, err := strconv.Atoi(val)
				if err != nil {
					return fmt.Errorf("invalid padding right")
				}
				cfg.PaddingRight = newPaddingRight
			case PADDING_BOTTOM:
				newPaddingBottom, err := strconv.Atoi(val)
				if err != nil {
					return fmt.Errorf("invalid padding bottom")
				}
				cfg.PaddingBottom = newPaddingBottom
			case PADDING_LEFT:
				newPaddingLeft, err := strconv.Atoi(val)
				if err != nil {
					return fmt.Errorf("invalid padding left")
				}
				cfg.PaddingLeft = newPaddingLeft
			}
		}
	}

	return nil
}

func GetConfig() *Config {
	return &cfg
}
