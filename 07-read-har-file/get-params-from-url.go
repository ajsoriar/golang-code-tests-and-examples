// example comes from: https://play.golang.org/p/mijE73rUgw
package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://www.test.com/url?foo=bar&foo=baz#this_is_fragment")
	fmt.Println("full uri:", u.String())
	fmt.Println("scheme:", u.Scheme)
	fmt.Println("opaque:", u.Opaque)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path", u.Path)
	fmt.Println("Fragment", u.Fragment)
	fmt.Println("RawQuery", u.RawQuery)
	fmt.Printf("query: %#v", u.Query())
}

// The response:

/*
full uri: http://www.test.com/url?foo=bar&foo=baz#this_is_fragment
scheme: http
opaque: 
Host: www.test.com
Path /url
Fragment this_is_fragment
RawQuery foo=bar&foo=baz
query: url.Values{"foo":[]string{"bar", "baz"}}
Program exited.
*/