package main

import (
  "net/http"
)

// @desc type Route is a struct, a typed collection of fields (a model).

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

// @desc type Routes is a slice, an ordered collection with contents of type Route

type Routes []Route

// @desc routes is a struct literal of Route types

const routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index
  },
  Route{
    "TodoIndex",
    "GET",
    "/todos",
    TodoIndex,
  },
  Route{
    "TodoShow",
    "GET",
    "/todos/{todoId}",
    TodoShow,
  },
  Route{
    "TodoCreate",
    "POST",
    "/todos",
    TodoCreate,
  },
}
