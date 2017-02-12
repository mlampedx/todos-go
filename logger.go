package main

import (
  "log"
  "net/http"
  "time"
)

// @desc Logger is a wrapper for handlers.go that prints the method, request uri, name, and elapsed time for each server request in the console.

func Logger(inner http.Handler, name string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    inner.ServeHTTP(w, r)
    log.Printf(
      "%s\t%s\t%s\t%s",
      r.Method,
      r.RequestURI,
      name,
      time.Since(start),
    )
  }) 
}
