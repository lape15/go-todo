package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	var todo Todo
	// id, _ := strconv.Atoi(req.URL.Query().Get("id"))
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	if id == -1 {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, v := range todos {
		if v.ID == id {
			todos[i] = todo
			fmt.Printf("Todos %+v\n", v)
		}
	}
	fmt.Print(len(todos))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated Todo with ID %d\n", id)
}
