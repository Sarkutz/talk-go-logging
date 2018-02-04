/* GoPro: Go Proxy */
package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	logrus "github.com/sirupsen/logrus"
)

type ctx struct {
	ReqId int64
}

const (
	hdr = "X-purl"

	peer_user_agent    = "ua"
	peer_origin_server = "origin"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	http.HandleFunc("/", proxyHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"pkg":           "main",
			logrus.ErrorKey: err,
		}).Panicln("Error in ListenAndServe")
	}
}

// proxyHandler is the main handler for HTTP requests
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := ctx{ReqId: time.Now().Unix()}
	logrus.WithFields(logrus.Fields{
		"pkg":     "main",
		"reqid":   ctx.ReqId,
		"peer":    peer_user_agent,
		"scheme":  getScheme(r.Header.Get(hdr)),
		"url":     r.Header.Get(hdr),
		"httpver": r.Proto,
		"method":  r.Method,
	}).Debugln("Received request from User Agent")

	// Convert user agent's request into request for origin server
	u := getURL(&ctx, r)
	if u == nil {
		logrus.WithFields(logrus.Fields{
			"pkg":   "main",
			"reqid": ctx.ReqId,
		}).Errorln("Error in getUrl")
		return
	}
	r.URL = u

	r.Header.Del(hdr)
	r.Host = r.URL.Host
	r.RequestURI = ""
	logrus.WithFields(logrus.Fields{
		"pkg":    "main",
		"reqid":  ctx.ReqId,
		"peer":   peer_origin_server,
		"scheme": getScheme(r.URL.Scheme),
		// "url":     r.URL,  // Print URL as JSON
		"url":     r.URL.String(),
		"httpver": r.Proto,
		"method":  r.Method,
	}).Debugln("Sending request to Origin Server")

	// Call origin server
	c := http.Client{}
	rs, err := c.Do(r)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"pkg":           "main",
			"reqid":         ctx.ReqId,
			logrus.ErrorKey: err,
		}).Errorln("Error in Client.Do", err)
		return
	}
	logrus.WithFields(logrus.Fields{
		"pkg":    "main",
		"reqid":  ctx.ReqId,
		"peer":   peer_origin_server,
		"scheme": getScheme(rs.Request.URL.Scheme),
		// "url":        rs.Request.URL,  // Print URL as JSON
		"url":        rs.Request.URL.String(),
		"httpver":    rs.Proto,
		"statuscode": rs.StatusCode,
	}).Debugln("Received response from Origin Server")

	// Proxy response to user agent
	w.WriteHeader(http.StatusOK)
	io.Copy(w, rs.Body)
}

// getURL returns the URL of the origin server
//
// The URL is obtained from the "X-purl" header in the user agent's request.
func getURL(ctx *ctx, r *http.Request) *url.URL {
	urlstr := r.Header.Get(hdr)
	u, err := url.Parse(urlstr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"pkg":           "main",
			"reqid":         ctx.ReqId,
			logrus.ErrorKey: err,
		}).Errorln("Error in url.Parse", err)
		return nil
	}
	return u
}

// getScheme returns the URL's scheme
func getScheme(url string) string {
	if url != "" {
		return strings.Split(url, ":")[0]
	}
	return ""
}
