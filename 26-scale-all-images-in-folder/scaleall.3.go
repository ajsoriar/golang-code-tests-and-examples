// ------------------------------------------
// Scale injecting new width as a parameter
// ------------------------------------------

/*
	EXAMPLE:

	> go run .\scaleall.3.go 1000

	Will scale all pictures in sources to width = 1000

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
	"strconv"

	"golang.org/x/image/draw"
)

func main() {

	var (
		rootFolder  string
		files       []string
		err         error
		fileName    string
		scaledWidth int
	)

	// 1) Console params
	arg := os.Args[1]
	fmt.Println("param: " + arg)

	scaledWidth, err = strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Argument is not a number!")
	}

	if InBetween(scaledWidth, 100, 1000) {
		fmt.Println("scaled width will be " + strconv.Itoa(scaledWidth))
	} else {
		scaledWidth = 100
		fmt.Println("Out of range! So scaled width will be 100.")
	}

	// 2) Get a list of all images in source folder

	rootFolder = "sources"

	files, err = IOReadDir(rootFolder)
	if err != nil {
		panic(err)
	}

	// 3) Does target directory exist?
	checkTargetDirectory()

	// 4) read the images one by one and scale

	for _, fileName = range files {

		fmt.Println(fileName)

		// 4.1) Open a new file to store the data of the image
		fl, err := os.Open(rootFolder + "/" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer fl.Close()

		// 4.2) Read image data and format
		sourceImage, imageFormat, err := image.Decode(fl)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(imageFormat)

		// 4.3) calculate new height
		scaledHeight := sourceImage.Bounds().Max.Y * scaledWidth / sourceImage.Bounds().Max.X

		// 4.4) Create new image canvas
		//newImage := image.Rect(0, 0, sourceImage.Bounds().Max.X/4, sourceImage.Bounds().Max.Y/4) // 4
		newImage := image.Rect(0, 0, scaledWidth, scaledHeight) // 4

		// 4.5) Scaling the image and fitting it to the new smaller canvas
		var res image.Image = scaleTo(sourceImage, newImage, draw.BiLinear)

		// bounds := sourceImage.Bounds()
		// y := bounds.Max.Y
		// x := bounds.Max.X
		// fmt.Println(x)
		// fmt.Println(y)

		// 4.6) Create the picture file
		file, err := os.Create("results/" + fileName)
		if err != nil {
			log.Fatal(err)
		}

		// 4.6) Encode picture data an store it in the file
		switch imageFormat {
		case "jpg":
		case "jpeg":

			// Specify the quality, between 0-100
			// Higher is better
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

		// 4.7) Close the file!
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

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
