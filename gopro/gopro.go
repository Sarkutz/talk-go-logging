/* GoPro: Go Proxy */
package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

const hdr = "X-purl"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("gopro:: ")

	http.HandleFunc("/", proxyHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln("Error in ListenAndServe", err)
	}
}

// proxyHandler is the main handler for HTTP requests
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("req = %+v\n", r)

	// Convert user agent's request into request for origin server
	u := getURL(r)
	if u == nil {
		log.Println("Error in getUrl")
		return
	}
	r.URL = u

	r.Header.Del(hdr)
	r.Host = r.URL.Host
	r.RequestURI = ""
	log.Printf("req = %+v\n", r)

	// Call origin server
	c := http.Client{}
	rs, err := c.Do(r)
	if err != nil {
		log.Println("Error in Client.Do", err)
		return
	}
	log.Printf("resp = %+v\n", rs)

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
	if err != nil {
		log.Println("Error in url.Parse", err)
		return nil
	}
	return u
}
