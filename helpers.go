package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createTodo(w http.ResponseWriter, req *http.Request) {
	var todo Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	fmt.Printf("Received Todo: %+v\n", todo)
	fmt.Print(nextID)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Received Todo: %+v\n", todo)
}

func getTodos(w http.ResponseWriter, req *http.Request) {

	jsonTodos, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonTodos)
}

func completeTodoHandler(w http.ResponseWriter, req *http.Request) {
	completed, err := json.Marshal(filterCompleted(todos))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(completed)
}

func filterCompleted(allTodos []Todo) []Todo {
	var completed []Todo
	for _, todo := range allTodos {
		if todo.Completed {
			completed = append(completed, todo)
		}
	}
	return completed
}

func editTodo(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	fmt.Print(id)
	w.WriteHeader(http.StatusOK)
}
