package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getHex(c string) string {

	hex := "null"

	switch c {
	case "0":
		hex = "x00"
	case "1":
		hex = "x01"
	case "2":
		hex = "x02"
	case "3":
		hex = "x03"
	case "4":
		hex = "x04"
	case "5":
		hex = "x05"
	case "6":
		hex = "x06"
	case "7":
		hex = "x07"
	case "8":
		hex = "x08"
	case "9":
		hex = "x09"
	}

	return hex
}

func main() {

	arg := os.Args[1]

	elements := []string{"\\x76"}

	//fmt.Println(arg)

	str := arg
	for _, r := range str {
		c := string(r)
		c = getHex(c)
		//fmt.Println(c)
		elements = append(elements, c)
	}

	fmt.Println(strings.Join(elements, "\\"))

	// echo -en '\x76\x01\x09\x02' > /dev/ttyAMA0
	// echo -en '\x78\x78\x78\x78' > /dev/ttyAMA0

	//

	// args := []string{"what", "ever", "you", "like"}
	// cmd := exec.Command(app, args...)

	cmd := exec.Command("echo -en '\x76\x01\x09\x02' > /dev/ttyAMA0")

	// Stdout buffer
	//cmdOutput := &bytes.Buffer{}

	// Attach buffer to command
	//cmd.Stdout = cmdOutput

	// Execute command
	//printCommand(cmd)
	//err :=
	cmd.Run() // will wait for command to return
	//printError(err)
	// Only output the commands stdout
	//printOutput(cmdOutput.Bytes()) // => go version go1.7.5 darwin/amd64

}

/*

// instructions

Build a binary with go build first.

$ go build command-line-arguments.go
$ ./command-line-arguments a b c d

Or just run it:

$ go run params.5.go 192

*/
