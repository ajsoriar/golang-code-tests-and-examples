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
)

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
		fmt.Println("Error: Num of files is diferent!")

		//Exit with status 3.
		os.Exit(3)

	} else {

		fmt.Println("All OK!")
	}

	// ------------------------------------------------------
	// 2) Create new file
	// ------------------------------------------------------

	newFile, err := os.Create("./gallery.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("/* ------- */")
	fmt.Println("/* Process */")
	fmt.Println("/* ------- */")

	// ------------------------------------------------------
	// 3) Process
	// ------------------------------------------------------

	var itemsString string = ""

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
		defer f1.Close()

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
		itemsString += compose_Item(cont, fileName)
		cont++
	}

	jsonData := compose_Gallery (itemsString) 

	// ------------------------------------------------------
	// 4) Write string in file
	// ------------------------------------------------------

	fmt.Println(jsonData)
	//newFile, err = os.Create("./gallery.json")
	newFile.Write([]byte (jsonData))

	// ------------------------------------------------------
	// 5) Close the file
	// ------------------------------------------------------

	newFile.Close()
	if err != nil {
		log.Fatal(err)
	}

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

func checkTargetDirectory(dirName string) { // Creates the folder if it doesn´t exist
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		log.Fatal(err)
		os.Exit(1)
	}
}

/*

var jsonData_Item = string (`
{
	"id": 1,
	"parent": 0,
	"type": "IMAGE",
	"description": "Picture long description",
	"zoom": 1,
	"name": "Picture name",
	"title": "Picture title",
	"tags":["A001", "A002", "A003"],
	"thumbnail": null,
	"target": null,
	"header": null,
	"footer": null,
	"background": null,
	"frame": null,
	"hover": null
}`)

*/

func compose_Item (cont int, fileName string) (string) {

	str := string (`
		{
			"id": `+ strconv.Itoa( cont ) +`,
			"parent": 0,
			"type": "IMAGE",
			"description": "Description of the picture `+ fileName +`",
			"zoom": 1,
			"name": "`+ fileName +`",
			"title": "`+ fileName +`",
			"tags":["ALL"],
			"thumbnail": null,
			"target": null,
			"header": null,
			"footer": null,
			"background": null,
			"frame": null,
			"hover": null
		}`)

	return str
}

/*

var jsonData_Gallery = string (`
{
    "galleryConfig": {
        "id": 202011041604498354406,
        "type": null,
        "name": "Name of the gallery",
        "title": "Title of the gallery",
        "description": "Description of the gallery",
        "background": {
            "color": "red",
            "video": {
                "src": null,
                "size": {
                    "w": 1024,
                    "h": 740
                }
            },
            "placeholder": null
        },
        "tags": [{
                "id": null,
                "label": "All"
            },
            {
                "id": 2,
                "label": "Smart TV"
            },
            {
                "id": 3,
                "label": "WEB"
            },
            {
                "id": 4,
                "label": "Demo"
            }
        ],
        "itemTypes": [],
        "debug": false
    },
    "items": []
}`)

*/

/*
var jsonData_Gallery_A = string (`
{
    "galleryConfig": {
        "id": 202011041604498354406,
        "type": null,
        "name": "Name of the gallery",
        "title": "Title of the gallery",
        "description": "Description of the gallery",
        "background": {
            "color": "red",
            "video": {
                "src": null,
                "size": {
                    "w": 1024,
                    "h": 740
                }
            },
            "placeholder": null
        },
        "tags": [{
                "id": null,
                "label": "All"
            },
            {
                "id": 2,
                "label": "Smart TV"
            },
            {
                "id": 3,
                "label": "WEB"
            },
            {
                "id": 4,
                "label": "Demo"
            }
        ],
        "itemTypes": [],
        "debug": false
    },
    "items": [
`)

var jsonData_Gallery_B = string (`]
}`)

*/

func compose_Gallery (items string) (string) {

str := string (`
{
	"galleryConfig": {
		"id": 202011041604498354406,
		"type": null,
		"name": "Name of the gallery",
		"title": "Title of the gallery",
		"description": "Description of the gallery",
		"background": {
			"color": "red",
			"video": {
				"src": null,
				"size": {
					"w": 1024,
					"h": 740
				}
			},
			"placeholder": null
		},
		"tags": [{
				"id": null,
				"label": "All"
			},
			{
				"id": 2,
				"label": "Smart TV"
			},
			{
				"id": 3,
				"label": "WEB"
			},
			{
				"id": 4,
				"label": "Demo"
			}
		],
		"itemTypes": [],
		"debug": false
	},
	"items": [`+ items +`
	]
}`)

	return str
}

