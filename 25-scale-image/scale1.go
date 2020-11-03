package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"time"

	"golang.org/x/image/draw"
)

const imgFile = "source.png"

// testing different scalers;

func main() {
	src := openImage()

	for _, scaller := range []struct {
		Name   string
		Scaler draw.Scaler
	}{
		{"NearestNeighbor", draw.NearestNeighbor},
		{"ApproxBiLinear", draw.ApproxBiLinear},
		{"BiLinear", draw.BiLinear},
		{"CatmullRom", draw.CatmullRom},
	} {

		newImage := image.Rect(0, 0, src.Bounds().Max.X/4, src.Bounds().Max.Y/4) // 4
		var res image.Image
		{ // show time to resize
			tp := time.Now()
			res = scaleTo(src, newImage, scaller.Scaler)
			log.Printf("scaling using %q takes %v time", scaller.Name, time.Now().Sub(tp))
		}

		file, err := os.Create(scaller.Name + "-demo.png")
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

// src   - source image
// rect  - size we want
// scale - scaler
func scaleTo(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	resultImg := image.NewRGBA(rect)
	scale.Scale(resultImg, rect, src, src.Bounds(), draw.Over, nil)
	return resultImg
}

/*
Results:

2020/11/03 11:49:13 scaling using "NearestNeighbor" takes 5.9187ms time
2020/11/03 11:49:13 scaling using "ApproxBiLinear" takes 11.0024ms time
2020/11/03 11:49:13 scaling using "BiLinear" takes 66.0083ms time
2020/11/03 11:49:13 scaling using "CatmullRom" takes 91.746ms time
*/
