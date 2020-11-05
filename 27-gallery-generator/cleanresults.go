package main

import (
	"log"
	"os"
)

// Deletes a folder from the current working directory

func main() {

	err := os.RemoveAll("results") // empty or not the folder is removed
	if err != nil {
		log.Fatal(err)
	}

	// err = os.Remove("results") // removes the folder when is empty. If not an error arises
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
