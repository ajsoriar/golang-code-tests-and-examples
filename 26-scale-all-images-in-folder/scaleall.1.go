package main

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/draw"
)

var imgFile = ""

func main() {

	var (
		rootFolder string
		files      []string
		err        error
		src        image.Image
		fileName   string
	)

	// 1) List all images in source

	rootFolder = "sources/"

	files, err = IOReadDir(rootFolder)
	if err != nil {
		panic(err)
	}

	for _, fileName = range files {
		fmt.Println(fileName)
	}

	// 2) read the images one by one and scale

	for _, fileName = range files {

		imgFile = rootFolder + fileName
		src = openImage()
		newImage := image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2) // 4

		var res image.Image = scaleTo(src, newImage, draw.BiLinear)

		file, err := os.Create("results/" + fileName + "-demo.png")
		if err != nil {
			log.Fatal(err)
		}

		err = png.Encode(file, res)
		file.Close()
		if err != nil {
			log.Fatal(err)
		}
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

func IOReadDir(rootFolder string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(rootFolder)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
