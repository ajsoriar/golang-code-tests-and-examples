package main

import "os"
import "fmt"

func main() {

	arg := os.Args[1]

	fmt.Println(arg)

	str := arg //"Hello"
	for _, r := range str {
		c := string(r)

		fmt.Println(c)
	}
}

/*

// instructions

Build a binary with go build first.

$ go build command-line-arguments.go
$ ./command-line-arguments a b c d

Or just run it:

$ go run params.2.go 192

*/
