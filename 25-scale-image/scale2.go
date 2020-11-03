package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

const imgFile = "source.png"

// testing different scalers;

func main() {
	src := openImage()

	newImage := image.Rect(0, 0, src.Bounds().Max.X/4, src.Bounds().Max.Y/4) // 4

	var res image.Image = scaleTo(src, newImage, draw.BiLinear)

	file, err := os.Create("result-demo.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(file, res)
	file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func openImage() image.Image {
	fl, err := os.Open(imgFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()
	img, _, err := image.Decode(fl)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func scaleTo(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	resultImg := image.NewRGBA(rect)
	scale.Scale(resultImg, rect, src, src.Bounds(), draw.Over, nil)
	return resultImg
}
