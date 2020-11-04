// This program renames all files in ./results folder adding .thumbnail at the end of the picture's name
// A. J. Soria R. 2020-11-04

// > go run .\rename.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

func main() {

	var (
		rootFolder string
		files      []string
		err        error
		oldName    string
		newName    string
		extension  string
		name       string
		cont       int
	)

	// 1) List all images in source

	rootFolder = "./results/"

	files, err = IOReadDir(rootFolder)
	if err != nil {
		panic(err)
	}

	cont = 0
	for _, oldName = range files {

		extension = path.Ext(oldName)
		name = oldName[0 : len(oldName)-len(extension)]
		newName = name + ".thumbnail" + extension
		err := os.Rename(rootFolder+oldName, rootFolder+newName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(oldName + " --> " + newName)
		cont++
	}

	fmt.Println(strconv.Itoa(cont) + " files were renamed!")
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
