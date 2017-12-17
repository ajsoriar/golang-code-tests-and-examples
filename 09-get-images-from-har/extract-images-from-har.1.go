package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "time"
    "net/url"
)

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

    file, e := ioutil.ReadFile("./demo.har")

    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }

    var arrOfData DataAndres
    json.Unmarshal(file, &arrOfData)
    //fmt.Printf("Results: %v\n", arrOfData )

    fmt.Printf("Results: %v\n", arrOfData.Log.Entries[1].Request.URL )
    
    for i := 0; i < len(arrOfData.Log.Entries) ; i++ {

        var myUrl string = arrOfData.Log.Entries[i].Request.URL;
        
        //fmt.Printf("Results: %v\n", myUrl)
        maprintUrlData( myUrl, i )
    }
}

func maprintUrlData( urlString string , num int) {

    //u, _ := url.Parse("http://www.test.com/url?foo=bar&foo=baz#this_is_fragment")
    u, _ := url.Parse( urlString  )
    fmt.Println("---", num,"---")
    fmt.Println("full uri:", u.String())
    fmt.Println("\n")
	fmt.Println("scheme:", u.Scheme)
	fmt.Println("opaque:", u.Opaque)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("Fragment:", u.Fragment)
	fmt.Println("RawQuery:", u.RawQuery)
    fmt.Printf("query: %#v", u.Query())
    fmt.Println("\n")
}