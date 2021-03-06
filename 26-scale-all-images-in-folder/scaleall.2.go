// ----------------------------------
// Scale percentage
// ----------------------------------

// ------------------------------------------
// Scale injecting new width as a parameter
// ------------------------------------------

/*
	EXAMPLE:

	> go run .\scaleall.2.go

	Will scale all pictures in sources.

	This way: width/4 and height/4

*/

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
		fileName   string
	)

	// 1) Get a list of all images in source folder

	rootFolder = "sources"

	files, err = IOReadDir(rootFolder)
	if err != nil {
		panic(err)
	}

	// 2) Does target directory exist?
	checkTargetDirectory()

	// 3) read the images one by one and scale

	for _, fileName = range files {

		fmt.Println(fileName)

		fileRoute = rootFolder + "/" + fileName

		// 3.1) Open a new file to store the data of the image
		fl, err := os.Open(fileRoute)
		if err != nil {
			log.Fatal(err)
		}
		defer fl.Close()

		// 3.2) Read image data and format
		sourceImage, imageFormat, err := image.Decode(fl)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(imageFormat)

		// 3.3) Create new image canvas
		newImage := image.Rect(0, 0, sourceImage.Bounds().Max.X/4, sourceImage.Bounds().Max.Y/4) // 4

		// 3.4) Scaling the image and fitting it to the new smaller canvas
		var res image.Image = scaleTo(sourceImage, newImage, draw.BiLinear)

		bounds := sourceImage.Bounds()
		y := bounds.Max.Y
		x := bounds.Max.X
		fmt.Println(x)
		fmt.Println(y)

		// 3.5) Create the picture file
		file, err := os.Create("results/scaled-" + fileName)
		if err != nil {
			log.Fatal(err)
		}

		// 3.6) Encode picture data an store it in the file
		switch imageFormat {
		case "jpg":
		case "jpeg":

			// Specify the quality, between 0-100, higher is better
			opt := jpeg.Options{
				Quality: 100,
			}

			err = jpeg.Encode(file, res, &opt)
			break

		case "gif":

			err = gif.Encode(file, res, nil)
			break

		case "png":

			err = png.Encode(file, res)
			break
		}

		// 3.7) Close the file!
		file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	// That was all!
}

func scaleTo(sourceImage image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	resultImg := image.NewRGBA(rect)
	scale.Scale(resultImg, rect, sourceImage, sourceImage.Bounds(), draw.Over, nil)
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
