package main

import (
    "image"
    "image/png"
    "image/color"
    //"math"
    "os"
)

func main() {

    var w, h int = 280, 240

    // Create a blan image 100x200 pixels
    myImage := image.NewRGBA(image.Rect(0, 0, w, h))
    for x := 0; x < w; x++ {
        for y := 0; y < h; y++ {
            c := color.RGBA{
                0,
                0,
                0,
                255,
            }
            myImage.Set(x, y, c)
        }
    }


    // outputFile is a File type which satisfies Writer interface
    outputFile, err := os.Create("test.png")
    if err != nil {
    	// Handle error
    }

    // Encode takes a writer interface and an image interface
    // We pass it the File and the RGBA
    png.Encode(outputFile, myImage)

    // Don't forget to close files
    outputFile.Close()
}
