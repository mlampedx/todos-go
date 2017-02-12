package main

import (
  "log"
  "net/http"
)

// @desc main declares a variable named router, assigns it the value of the output of the NewRouter function, and listens on port 8080.

func main() {
  router := NewRouter()
  log.Fatal(http.ListenAndServe(":8080", router))
}
