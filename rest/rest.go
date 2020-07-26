// Scope: REST
package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"

	log "github.com/sirupsen/logrus"
)

// MimeType is a collection of mime types
var MimeType = struct {
	Json, Html, Text string
}{
	"application/json", "text/html", "text/plain",
}

// RenderJsonErr...
func RenderJsonErr(resp http.ResponseWriter, httpStatus int, err error) {
	resp.Header().Set("Content-Type", MimeType.Json)
	resp.WriteHeader(httpStatus)
	errStr := "bad request"
	if err != nil {
		errStr = err.Error()
	}
	jsonErr := struct {
		Error string `json:"error"`
	}{
		errStr,
	}
	json.NewEncoder(resp).Encode(&jsonErr)
}

// HttpRedirectTo...
func HttpRedirectTo(resp *http.ResponseWriter, location string) {
	(*resp).Header().Set("Location", location)
	(*resp).WriteHeader(302)
}

func HttpChainMiddlewareRequestDump(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		log.WithField("middleware", "request_dump").Info(req.Method)
		log.WithField("middleware", "request_dump").Info(req.RemoteAddr)
		log.WithField("middleware", "request_dump").Info(req.URL.Path)
		log.WithField("middleware", "request_dump").Info(req.URL.RawQuery)
		log.WithField("middleware", "request_dump").Info(req.Header.Get("Accept-Encoding"))
		log.WithField("middleware", "request_dump").Info(req.Header["User-Agent"])
		log.WithField("middleware", "request_dump").Info(req.ContentLength)
		log.Info(req.Header)
		reqDump, _ := httputil.DumpRequest(req, true)
		log.WithField("middleware", "request_dump").Info(string(reqDump))
		next(resp, req)
	}
}
