package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/draw"
)

var fileRoute = ""

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

	// 2) read the images one by one and scale

	for _, fileName = range files {

		fmt.Println(fileName)

		fileRoute = rootFolder + fileName

		fl, err := os.Open(fileRoute)
		if err != nil {
			log.Fatal(err)
		}
		defer fl.Close()
		img, imageFormat, err := image.Decode(fl)
		if err != nil {
			log.Fatal(err)
		}

		src = img

		fmt.Println(imageFormat)

		newImage := image.Rect(0, 0, src.Bounds().Max.X/4, src.Bounds().Max.Y/4) // 4

		var res image.Image = scaleTo(src, newImage, draw.BiLinear)

		bounds := src.Bounds()
		y := bounds.Max.Y
		x := bounds.Max.X
		fmt.Println(x)
		fmt.Println(y)

		checkTargetDirectory()

		file, err := os.Create("results/" + fileName + "-demo." + imageFormat)
		if err != nil {
			log.Fatal(err)
		}

		switch imageFormat {
		case "jpg":
			err = jpeg.Encode(file, res, nil)
			break

		case "gif":

			err = gif.Encode(file, res, nil)
			break

		case "png":
			err = png.Encode(file, res)
			break
		}

		file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
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

func checkTargetDirectory() { // Creates the folder if it doesn´t exist
	_, err := os.Stat("results")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("results", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
