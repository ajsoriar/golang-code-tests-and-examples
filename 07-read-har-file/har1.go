package main

/*
cat config.json
{"object": 
    {
       "buffer_size": 10,
       "Databases":
       [
               {
                       "host": "localhost",
                       "user": "root",
                       "pass": "",
                       "type": "mysql",
                       "name": "go",
                       "Tables":
                       [
                               {
                                       "name": "testing",
                                       "statment": "teststring",
                                       "regex": "teststring ([0-9]+) ([A-z]+)",
                                       "Types": 
                                        [
                                           {
                                               "id": "int",
                                               "value": "string"
                                           }
                                        ]
                               }
                       ]
               }
       ]
    }
}

*/

import (
    "fmt"
    "os"
    //"json"
    "encoding/json"
    "io/ioutil"
)

/*

type jsonobject struct {
    Object ObjectType
}

type ObjectType struct {
    Buffer_size int
    Databases   []DatabasesType
}

type DatabasesType struct {
    Host   string
    User   string
    Pass   string
    Type   string
    Name   string
    Tables []TablesType
}

type TablesType struct {
    Name     string
    Statment string
    Regex    string
    Types    []TypesType
}

type TypesType struct {
    Id    string
    Value string
}

*/


// --------------------

/*
{
    "log": {
      "version": "1.2",
      "creator": {
        "name": "WebInspector",
        "version": "537.36"
      },
      "pages": [
        {
          "startedDateTime": "2017-12-08T10:30:51.319Z",
          "id": "page_1",
          "title": "https://codereview.stackexchange.com/questions/182306/javascript-improve-a-sequence-of-if-conditions",
          "pageTimings": {
            "onContentLoad": 13706.380000003264,
            "onLoad": 27442.97999999253
          }
        }
      ],
      "entries": [
*/

type jsonobject struct {
    ObjectAndresLog LogType
}

type LogType struct {
    //Buffer_size int
    //Databases   []DatabasesType
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
    //file, e := ioutil.ReadFile("./config.json")
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