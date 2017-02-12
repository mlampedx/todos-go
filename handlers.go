package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

// @desc Index handler prints "Welcome" when a user visits the "/" route.

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}

// @desc TodoIndex handler sets header and sends 200 status code. 
// @err If there is an error, print a JSON encoding of the todos error to the stream.

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(todos);
  err != nil {
    panic(err)
  }
}

// @desc TodoShow declares the variable routeMap and assigns it the value of a map of route variables. We find the todo located at the todoId param and print it. 

func TodoShow(w http.ResponseWriter, r *http.Request) {
  routeMap := mux.Vars(r)
  todoId := routeMap["todoId"]
  fmt.Fprintln(w, "Todo show:", todoId)
}

// @desc TodoCreate limits the size of the request body, handles errors, and sends a status code of 201 and creates a todo in the mock db, if successful. 
// @err1 If there is an error when reading the body, send the error.
// @err2 If there is an error when the client closes the response body, send the error.
// @err3 If there is an error when the body is unmarshalled to the Todo struct, send a status code of 422 and send the error.

func TodoCreate(w http.ResponseWriter, r *http.Request) {
  var todo Todo
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  
  if err != nil {
    panic(err)
  }

  if err := r.Body.Close();
  err != nil {
    panic(err)
  }

  if err := json.Unmarshal(body, &todo);
  err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err);
    err != nil {
      panic(err)
    }
  }

  t := MockDBCreateTodo(todo)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(t);
  err != nil {
    panic(err)
  }
}