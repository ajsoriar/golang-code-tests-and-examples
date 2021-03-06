package main

import (
    "fmt"
    "os"
    //"json"
    "encoding/json"
    "io/ioutil"
)

type jsonobject struct {
    ObjectAndresLog LogType
}

type LogType struct {
    version int
    creator CreatorType
    pages PagesType
    entries EntriesType
}

type CreatorType struct {
    name string
    version string
}

type PagesType struct {
    startedDateTime string //": "2017-12-08T10:30:51.319Z",
    id string // "page_1",
    title string //https://codereview.stackexchange.com/questions/182306/javascript-improve-a-sequence-of-if-conditions",
    //pageTimings": {
}

type EntriesType struct {
    //name string
    //version string
}

func main() {
    file, e := ioutil.ReadFile("./demo.har")

    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    //m := new(Dispatch)
    //var m interface{}
    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results: %v\n", jsontype)
}