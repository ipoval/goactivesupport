// Scope: REST
package goactivesupport

import (
  "net/http"
  "net/http/httputil"

  log "github.com/sirupsen/logrus"
)

var MimeType = struct {
  Json, Html, Text string
}{
  "application/json", "text/html", "text/plain",
}

func HttpRedirectTo(resp *http.ResponseWriter, location string) {
  (*resp).Header().Set("Location", location)
  (*resp).WriteHeader(302)
}

func HttpChainMiddlewareRequestDump(next http.HandlerFunc) http.HandlerFunc {
  return func(resp http.ResponseWriter, req *http.Request) {
    log.WithField("middleware", "request_dump").Info(req.Method)
    log.WithField("middleware", "request_dump").Info(req.URL.Path)
    log.WithField("middleware", "request_dump").Info(req.URL.Path[1:])
    log.WithField("middleware", "request_dump").Info(req.RemoteAddr)
    log.WithField("middleware", "request_dump").Info(req.Header.Get("Accept-Encoding"))
    log.WithField("middleware", "request_dump").Info(req.Header["Content-Type"])
    log.WithField("middleware", "request_dump").Info(req.ContentLength)
    log.Info(req.Header)
    reqDump, _ := httputil.DumpRequest(req, true)
    log.WithField("middleware", "request_dump").Info(string(reqDump))
    next(resp, req)
  }
}
