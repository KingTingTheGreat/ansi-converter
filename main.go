package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/kingtingthegreat/ansi-converter/config"
	"github.com/kingtingthegreat/ansi-converter/converter"
)

func main() {
	err := config.ParseArgs()
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := config.GetConfig()

	file, err := os.Open(cfg.FilePath)
	if err != nil {
		log.Fatal("error opening file")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("error decoding image")
	}

	x := converter.Convert(img, cfg)
	fmt.Println(x)
}
