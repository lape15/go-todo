package main

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/todos", todoHandler)
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":5000", nil)

}

func todoHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		getTodos(w, req)
	case http.MethodPost:
		createTodo(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
