package main

import (
	"fmt"
)

var currentId int
var todos Todos

func init(){
	RepoCreateTodo(Todo{Name: "Learn Golang!"})
	RepoCreateTodo(Todo{Name: "Find remote!"})
	RepoCreateTodo(Todo{Name: "Move to Japan!"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos{
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo{
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyID(id int) error{
	for i, t := range todos{
		if t.Id == id {
			todos = append(todos[:i], todos[:i+1]...)
			return nil
		}
	}
	return fmt.Errorf("Couldnt find Todo with id %d  in database", id)
}