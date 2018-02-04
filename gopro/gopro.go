/* GoPro: Go Proxy */
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const hdr = "X-purl"

func main() {
	http.HandleFunc("/", proxyHandler)

	err := http.ListenAndServe(":8080", nil)
	checkError(err)
}

// proxyHandler is the main handler for HTTP requests
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("req = %+v\n", r)

	// Convert user agent's request into request for origin server
	r.URL = getURL(r)
	r.Header.Del(hdr)
	r.Host = r.URL.Host
	r.RequestURI = ""
	fmt.Printf("req = %+v\n", r)

	// Call origin server
	c := http.Client{}
	rs, err := c.Do(r)
	checkError(err)
	fmt.Printf("resp = %+v\n", rs)

	// Proxy response to user agent
	w.WriteHeader(http.StatusOK)
	io.Copy(w, rs.Body)
}

// getURL returns the URL of the origin server
//
// The URL is obtained from the "X-purl" header in the user agent's request.
func getURL(r *http.Request) *url.URL {
	urlstr := r.Header.Get(hdr)
	u, err := url.Parse(urlstr)
	checkError(err)
	return u
}

// CheckError panics on error
//
// This is useful for PoC, but not suitable for production.
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
