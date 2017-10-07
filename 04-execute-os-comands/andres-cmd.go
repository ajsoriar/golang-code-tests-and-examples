/*

- Use this example this way: 

    http://localhost:8000/andres

- this example will execute an unix command in the server

	
*/

package main

import (
    //"io"
    "net/http"
    "fmt"
    "os"
    //"encoding/json"
    //"io/ioutil"
    "os/exec"
    "strings"
    "bytes"
)

var counter int

func main() {

	fmt.Printf("%s\n", "Server will be started in http://localhost:8000/" )

    fmt.Printf("%s\n", "main()" )
	http.HandleFunc("/andres", helloAndres )

	http.ListenAndServe(":8000", nil) // Starts a server!

    counter++
}

// I have defined the following functions to keep the examples short:

func printCommand(cmd *exec.Cmd) {
  fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
  if err != nil {
    os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
  }
}

func printOutput(outs []byte) {
  if len(outs) > 0 {
    fmt.Printf("==> Output: %s\n", string(outs))
  }
}

func helloAndres(w http.ResponseWriter, r *http.Request) {

	// The first and most obvious use is to collect output from an external command. 
	// An easy way to do that is to use the CombinedOutput function.

	/*
		// Create an *exec.Cmd
		cmd := exec.Command("echo", "Called from Go!")

		// Combine stdout and stderr
		printCommand(cmd)
		output, err := cmd.CombinedOutput()
		printError(err)
		printOutput(output) // => go version go1.3 darwin/amd64
    */

	// This works well if you also want to check for any error messages output but if you 
	// want finer control over the output of a command then we can route it into different 
	// buffers giving us control over both standard output and standard error.


	// Create an *exec.Cmd
	//cmd := exec.Command("go", "version")
	cmd := exec.Command("ls", "-l")

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}

	// Attach buffer to command
	cmd.Stdout = cmdOutput

	// Execute command
	printCommand(cmd)
	err := cmd.Run() // will wait for command to return
	printError(err)
	// Only output the commands stdout
	printOutput(cmdOutput.Bytes()) // => go version go1.7.5 darwin/amd64

}
