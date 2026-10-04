package main

import (
	"todo-list/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	h := handler.HandlerUserConstructor()
	t := handler.HandlerTaskConsctructor()

	r.HandleFunc("/users", h.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", h.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", h.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", h.ChangeUserHandler).Methods("PATCH")
	r.HandleFunc("/users/{id}", h.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/users/{id}/tasks", t.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", t.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", t.GetAllTaskHandler).Methods("GET")
	r.HandleFunc("/users/{id}/tasks", t.GetUserTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", t.ChangeTaskHandler).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", t.DeleteTaskHandler).Methods("DELETE").	
}
