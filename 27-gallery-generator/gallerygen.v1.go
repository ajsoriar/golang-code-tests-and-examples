// ------------------------------------------
// Generate gallery data 
// ------------------------------------------

package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

var galleryFolderName = "golang-generated-gallery"
var sourcesDataPath = "./galleries"

func main() {

	var (
		rootFolder string
		files      []string
		err        error
		fileName   string
	)

	// ------------------------------------------------------
	// 1) Count the number of files in source and thumbnails
	// ------------------------------------------------------

	fmt.Println("/* ----- */")
	fmt.Println("/* check */")
	fmt.Println("/* ----- */")

	cont1 := 0
	cont2 := 0

	// A)

	rootFolder = "results"

	checkTargetDirectory(rootFolder)

	files, err = IOReadDir(rootFolder)
	if err != nil {
		panic(err)
	}

	for _, fileName = range files {
		fmt.Println(fileName)
		cont2++
	}

	// B)

	rootFolder = "sources"

	checkTargetDirectory(rootFolder)

	files, err = IOReadDir(rootFolder)
	if err != nil {
		panic(err)
	}

	for _, fileName = range files {
		fmt.Println(fileName)
		cont1++
	}

	if cont1 != cont2 {

		defer fmt.Println("!") // INTERESTING: defers will not be run when using os.Exit, so this fmt.Println will never be called.
		fmt.Println("Error: Num of files is different!")

		//Exit with status 3.
		os.Exit(3)

	} else {

		fmt.Println("All OK!")
	}

	// ------------------------------------------------------
	// 2) Create new file
	// ------------------------------------------------------

	newFile, err := os.Create(galleryFolderName +".json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("/* ------- */")
	fmt.Println("/* Process */")
	fmt.Println("/* ------- */")

	// ------------------------------------------------------
	// 3) Process
	// ------------------------------------------------------

	var itemsString = ""

	cont := 0

	for _, fileName = range files {

		var route = "./" + rootFolder + "/" + fileName
		fmt.Println(route)

		// 3.1) Open, load into memory and close the image

		f1, err := os.Open(route)
		if err != nil {
			fmt.Println("Err 1")
			log.Fatal(err)
		}
		defer f1.Close() // TODO: Why I just can't remove 'defer'?

		// 3.2) Decode the picture: Read image data and format

		sourceImage, imageFormat, err := image.Decode(f1)
		if err != nil {
			fmt.Println("Err 2")
			log.Fatal(err)
		}

		// 3.3) Extract info from the picture

		fmt.Println(" - imageFormat: " + imageFormat)
		bounds := sourceImage.Bounds()
		h := bounds.Max.Y
		w := bounds.Max.X
		fmt.Println(" - x: " + strconv.Itoa(w))
		fmt.Println(" - y: " + strconv.Itoa(h))

		// 3.4) Concat strings to create a json

		if cont > 0 { itemsString += "," }

		itemsString += composeItem(cont, fileName, sourcesDataPath +"/"+ galleryFolderName +"/images/", w, h, "#e8e8e8")
		cont++
	}

	jsonData := composeGallery(itemsString)

	// ------------------------------------------------------
	// 4) Write string in file
	// ------------------------------------------------------

	fmt.Println(jsonData)
	newFile.Write([]byte (jsonData))

	// ------------------------------------------------------
	// 5) Close the file
	// ------------------------------------------------------

	newFile.Close()

	// That was all!
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

func checkTargetDirectory(dirName string) {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		log.Fatal(err)
		os.Exit(1)
	}
}

func composeItem(
	cont int, 
	fileName string,
	filePath string,
	pictureWidth int,
	pictureHeight int,
	pictureColor string ) string {

	str := `
		{
			"id": `+ strconv.Itoa( cont ) +`,
			"parent": 0,
			"type": "IMAGE",
			"description": "Description of the picture `+ fileName +`",
			"zoom": 1,
			"name": "`+ fileName +`",
			"title": "`+ fileName +`",
			"tags":["ALL"],
			"thumbnail": {
                "size":{
                    "w": `+ strconv.Itoa( pictureWidth ) +`,
                    "h": `+ strconv.Itoa( pictureHeight ) +`
                },
				"src": "`+ filePath + fileName +`",
				"placeholderColor": "`+ pictureColor +`"
            },
			"target": {
                "size":{
                    "w": `+ strconv.Itoa( pictureWidth ) +`,
                    "h": `+ strconv.Itoa( pictureHeight ) +`
                },
				"src": "`+ filePath + fileName +`"
			},
			"background": {
				"color": "`+ pictureColor +`",
				"image": null,
				"video": null
			},
			"header": null,
			"footer": null,
			"frame": null,
			"hover": {
				"border": false,
				"zoom": true,
				"translucent": false,
				"overlay": true,
				"banner": false,
				"shadow": false,
				"overlayText": false
			}
		}`

	return str
}

func composeGallery(items string) string {

	t := time.Now()
	tUnixNano := t.UnixNano()
	timestamp := strconv.Itoa(int(tUnixNano)) 
	galleryName := "Unnamed Gallery"
	galleryTitle := "Untitled Gallery"
	galleryDescription := "Description of this gallery"
	galleryBackgroundColor := "#ffffff"

str := `
{
	"galleryConfig": {
		"id": `+ timestamp +`,
		"type": null,
		"name": "`+ galleryName +`",
		"title": "`+ galleryTitle +`",
		"description": "`+ galleryDescription +`",
		"background": {
			"color": "`+ galleryBackgroundColor +`",
			"gradient": {
                "color1": "`+ galleryBackgroundColor +`",
                "color2": "`+ galleryBackgroundColor +`",
                "angle": null
			},
			"video": null,
			"placeholder": null
		},
		"tags": [{
				"id": null,
				"label": "All"
			}
		],
		"itemTypes": [],
		"debug": false
	},
	"items": [`+ items +`
	]
}`

	return str
}

