package main

import (
	"net/http"
	//"path"
	//"fmt"
	"bytes"

    "image"
    "image/png"
    "image/color"
    //"math"
	//"os"
)

func main() {

	http.HandleFunc("/", _welcome)
	http.HandleFunc("/getImage", _image_request)

		// getImage, params
		// url?href=/

	http.ListenAndServe(":7009", nil) // Check out this to stop the server: https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve
}
// -----------------------------------------------
// - WELCOME
// -----------------------------------------------

//http://localhost:7009/
func _welcome(w http.ResponseWriter, r *http.Request) {
	js := "<b>Hello HTML! This is a go server!</b>"
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Server", "A Go Web Server")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(js))
}

// -----------------------------------------------
// - WELCOME
// -----------------------------------------------

func _image_request (w http.ResponseWriter, r *http.Request) {

	//fp := "hello!"
	
/*
	switch r.FormValue("width") {
	case "ATPFormat":
		//esc = "&amp;"
		fp = path.Join("data/andres", "response_1.json")'
		_ = fp // Get rid of: fp declared and not used

	default:
		fp = path.Join("data/andres", "response_2.json")
		_ = fp // Get rid of: fp declared and not used
	}
*/
	//var default_w, default_h int = 640, 480
	//var format = "png"
	//var bg_color = color.RGBA{255,0,0,255} // red
	//var file_name = "test"
	//base64 = false

	//url_w := r.FormValue("width")
	//url_h := r.FormValue("height")
	//url_format := r.FormValue("format")
	//url_file_name := r.FormValue("filename")

	// set values
	img_w := 300 //url_w
	img_h := 200 //url_h
    
    // Create a blank image 100x200 pixels
    myImage := image.NewRGBA(image.Rect(0, 0, img_w, img_h))
    for x := 0; x < img_w; x++ {
        for y := 0; y < img_h; y++ {
            c := color.RGBA{ 50, 50, 50, 255 }
            myImage.Set(x, y, c)
        }
	}
	myImage.Set(0, 0, color.RGBA{255, 0, 0, 255})
    myImage.Set(img_w -1, 0, color.RGBA{255, 0, 0, 255})
    myImage.Set(img_w -1, img_h -1, color.RGBA{255, 0, 0, 255})
    myImage.Set(0, img_h -1, color.RGBA{255, 0, 0, 255})

	/*

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
	
	*/

	buffer := new(bytes.Buffer)

	if err := png.Encode(buffer, myImage); err != nil {
		//f.Close()
		//log.Fatal(err)
	}

	//fp := path.Join("data/andres", "response_1.json")
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Server", "A Go Web Server")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// buffer.Bytes()
	//http.ServeFile(w, r, fp)
	//w.Write([]byte( buffer ))
	w.Write( buffer.Bytes() )
}
