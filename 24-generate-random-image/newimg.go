package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var mySeed int64 = 1

func main() {

	// 1) random width and height from 0 yo 4K
	imgW := randomInt(100, 3840)
	imgH := randomInt(100, 3840)

	// 2) Create the image
	myImage := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	r := uint8(randomInt(0, 255))
	g := uint8(randomInt(0, 255))
	b := uint8(randomInt(0, 255))
	for x := 0; x < imgW; x++ {
		for y := 0; y < imgH; y++ {
			c := color.RGBA{r, g, b, 255}
			myImage.Set(x, y, c)
		}
	}

	// 3) Paint pixels on each corner
	myImage.Set(0, 0, color.RGBA{255, 0, 0, 255})
	myImage.Set(imgW-1, 0, color.RGBA{255, 0, 0, 255})
	myImage.Set(imgW-1, imgH-1, color.RGBA{255, 0, 0, 255})
	myImage.Set(0, imgH-1, color.RGBA{255, 0, 0, 255})

	// 4) Generate the name of the file

	t := time.Now()
	tUnixNano := t.UnixNano()
	fileName := strconv.Itoa(int(tUnixNano)) + "_" + strconv.Itoa(int(imgW)) + "_" + strconv.Itoa(int(imgH)) + ".png"

	// 5) Create a file
	outputFile, err := os.Create(fileName)
	if err != nil {
		// Handle error
	}

	// 6) Encode takes a writer interface and an image interface
	png.Encode(outputFile, myImage)

	// 7) Close file!
	outputFile.Close()
}

func randomInt(min, max int) (result int) {
	//rand.Seed(time.Now().UnixNano()) // Sory Andres this is just not enough! LOL ( So I will use a small trick: mySeed )
	mySeed++
	rand.Seed(time.Now().UnixNano() + mySeed)
	result = rand.Intn(max-min+1) + min
	return result
}
