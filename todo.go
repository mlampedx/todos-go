package main

import "time"

// @desc type Todo is a struct, a typed collection of fields (a model).

type Todo struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Completed bool `json:"completed"`
  Due time.Time `json:"due"`
}

// @desc type Todos is a slice, an ordered collection with contents of type Route

type Todos []Todo
