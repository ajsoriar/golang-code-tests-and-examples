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

    // Create a black image
    myImage := image.NewRGBA( image.Rect(0, 0, w, h) )
    for x := 0; x < w; x++ {
        
        for y := 0; y < h; y++ {

            c := color.RGBA{ 0, 0, 0, 255, }
            myImage.Set(x, y, c)
        }
    }

    myImage = vLine( 0, 100, color, myImage) // Draw vertical line

    myImage = hLine( 0, 100, color, myImage) // Draw horizontal line

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

/*
func pset(x, y) {

}
*/

func vLine(y, lon int, c color ,img image ) {

	for pos = y; y < (y + lon); y++ {

        img = myImage.Set(x, y, c)

        pset(x, y) 
    }
    
    return img
}

func hLine( x, lon int, c color ,img image ) {

	for pos = x; x < (x + lon); x++ {

        img = myImage.Set(x, y, c)

        pset(x, y) 
    }
    
    return img
}

func drawLine(sx, sy, ex, ey int) {

	for x = sx; x != ex; x += inc {

		if sum > w {

            pset(x, y) 
		}
	}
}