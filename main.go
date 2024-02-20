package main

import (
	"fmt"
	"net/http"

	// "gorilla/mux"
	"github.com/gorilla/mux"
)

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	ID        int    `json:"id"`
}

var (
	todos  []Todo
	nextID int
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todo/{id}", singleTodoHandler)
	r.HandleFunc("/todos", todosHandler)
	r.HandleFunc("/", handleRequest)
	// http.Handle("/", r)
	http.ListenAndServe(":5000", r)

}

func singleTodoHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Crap")
	switch req.Method {
	case http.MethodPatch:
		editTodo(w, req)
	default:
		fmt.Print(http.MethodPatch)
		http.Error(w, "Method not allowed there", http.StatusMethodNotAllowed)
	}
}

func todosHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getTodos(w, req)
	case http.MethodPost:
		createTodo(w, req)
	default:
		http.Error(w, "Method not allowed here", http.StatusMethodNotAllowed)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is a GET request!")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is a POST request!")
}
