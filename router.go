package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

// @desc NewRouter uses the mux HTTP request multiplexer package to create a router.
// @return A new router with logger-wrapped routes that dictate allowed methods, path, name, and handler based on the Route struct.

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    
    var handler http.Handler
    handler = route.HandlerFunc
    handler = Logger(handler, route.Name)

    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(route.HandlerFunc)
  }

  return router
}
