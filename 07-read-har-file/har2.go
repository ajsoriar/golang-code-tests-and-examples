package main

import (
    "fmt"
    "os"
    //"json"
    "encoding/json"
    "io/ioutil"
    "time"
)

/*
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

*/

//This struct was created thanks to: https://mholt.github.io/json-to-go/

type DataAndres struct {
    Log struct {
        Version string `json:"version"`
        Creator struct {
            Name    string `json:"name"`
            Version string `json:"version"`
        } `json:"creator"`
        Pages []struct {
            StartedDateTime time.Time `json:"startedDateTime"`
            ID              string    `json:"id"`
            Title           string    `json:"title"`
            PageTimings     struct {
                OnContentLoad float64 `json:"onContentLoad"`
                OnLoad        float64 `json:"onLoad"`
            } `json:"pageTimings"`
        } `json:"pages"`
        Entries []struct {
            StartedDateTime time.Time `json:"startedDateTime"`
            Time            float64   `json:"time"`
            Request         struct {
                Method      string `json:"method"`
                URL         string `json:"url"`
                HTTPVersion string `json:"httpVersion"`
                Headers     []struct {
                    Name  string `json:"name"`
                    Value string `json:"value"`
                } `json:"headers"`
                QueryString []interface{} `json:"queryString"`
                Cookies     []interface{} `json:"cookies"`
                HeadersSize int           `json:"headersSize"`
                BodySize    int           `json:"bodySize"`
            } `json:"request"`
            Response struct {
                Status      int    `json:"status"`
                StatusText  string `json:"statusText"`
                HTTPVersion string `json:"httpVersion"`
                Headers     []struct {
                    Name  string `json:"name"`
                    Value string `json:"value"`
                } `json:"headers"`
                Cookies []struct {
                    Name     string    `json:"name"`
                    Value    string    `json:"value"`
                    Path     string    `json:"path"`
                    Domain   string    `json:"domain"`
                    Expires  time.Time `json:"expires"`
                    HTTPOnly bool      `json:"httpOnly"`
                    Secure   bool      `json:"secure"`
                } `json:"cookies"`
                Content struct {
                    Size     int    `json:"size"`
                    MimeType string `json:"mimeType"`
                    Text     string `json:"text"`
                } `json:"content"`
                RedirectURL  string `json:"redirectURL"`
                HeadersSize  int    `json:"headersSize"`
                BodySize     int    `json:"bodySize"`
                TransferSize int    `json:"_transferSize"`
            } `json:"response"`
            Cache struct {
            } `json:"cache"`
            Timings struct {
                Blocked         float64 `json:"blocked"`
                DNS             int     `json:"dns"`
                Ssl             int     `json:"ssl"`
                Connect         int     `json:"connect"`
                Send            float64 `json:"send"`
                Wait            float64 `json:"wait"`
                Receive         float64 `json:"receive"`
                BlockedQueueing int     `json:"_blocked_queueing"`
            } `json:"timings"`
            ServerIPAddress string `json:"serverIPAddress"`
            Connection      string `json:"connection,omitempty"`
            Pageref         string `json:"pageref"`
            FromCache       string `json:"_fromCache,omitempty"`
        } `json:"entries"`
    } `json:"log"`
}

func main() {
    //file, e := ioutil.ReadFile("./config.json")
    file, e := ioutil.ReadFile("./demo.har")

    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    // fmt.Printf("%s\n", string(file))

    //m := new(Dispatch)
    //var m interface{}

    /*
    var jsontype jsonobject
    json.Unmarshal(file, &jsontype)
    fmt.Printf("Results: %v\n", jsontype)
    */

    var arrOfData DataAndres
    json.Unmarshal(file, &arrOfData)
    //fmt.Printf("Results: %v\n", arrOfData )

    fmt.Printf("Results: %v\n", arrOfData.Log.Entries[1].Request.URL )
    
    for i := 0; i < len(arrOfData.Log.Entries) ; i++ {

        fmt.Printf("Results: %v\n", arrOfData.Log.Entries[i].Request.URL )
    }
}

/*
func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}
*/