package main

import "fmt"

// @desc Declaring variables and assigning them types

var currentId int
var todos Todos

// @desc init arbitrarily creates two todos for the mockdb

func init() {
  MockDBCreateTodo(Todo{Name: "Task One"})
  MockDBCreateTodo(Todo{Name: "Task Two"})
}

// @desc MockDBFindTodo uses the range function to iterate over the slice of todos seeking a todo with a matching integer id to the argument passed.
// @return If successful, returns the todo with a matching id. If unsuccessful, returns an empty todo.

func MockDBFindTodo(id int) Todo {
  for _, t := range todos {
    if t.Id == id {
      return t
    }
  }
  return Todo{}
}

// @desc MockDBCreateTodo appends a new todo to the todos slice with the content passed as an argument.
// @return The newly created todo.

func MockDBCreateTodo(t Todo) Todo {
  currentId += 1
  t.Id = currentId
  todos = append(todos, t)
  return t
}

// @desc MockDBDeleteTodo uses the range function to iterate over the slice of todos seeking a todo with a matching integer id to the argument passed.
// @return If successful, removes the todo with a matching id and then combines two slices into one slice. If unsuccessful, prints error message.

func MockDBDeleteTodo(id int) error {
  for i, t := range todos {
    if t.Id == id {
      todos = append(todos[:i], todos[i+1:]...)
      return nil
    }
  }
  return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
