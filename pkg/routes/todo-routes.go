package routes

import (
	"github.com/gorilla/mux"
	"github.com/ritik-agarwal50/go_api_tutorial/pkg/controllers"
)

var RegisterTodoRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/api/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", controllers.GetTodoById).Methods("GET")
	router.HandleFunc("/api/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/api/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

}
