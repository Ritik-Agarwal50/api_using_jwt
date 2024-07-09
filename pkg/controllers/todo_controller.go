package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ritik-agarwal50/go_api_tutorial/pkg/models"
	"github.com/ritik-agarwal50/go_api_tutorial/pkg/utils"
)

var NewTodo models.Todo

func GetTodos(w http.ResponseWriter, r *http.Request) {
	NewTodo := models.GetAllTodos()
	res, _ := json.Marshal(NewTodo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	createTodo := &models.Todo{}
	utils.ParseBody(r, createTodo)
	createTodo.CreateTodo()
	res, _ := json.Marshal(createTodo)
	w.Header().Set("Cintent-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

}
func GetTodoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoId := params["id"]
	Id, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		return
	}
	todo, _ := models.GetTodoById(Id)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	//fetch the all the data of respective id
	//then we will delete the entire data of that id
	//then we will create the new data with the same id
	var updatetodo = &models.Todo{}
	utils.ParseBody(r, updatetodo)
	params := mux.Vars(r)
	todoId := params["id"]
	Id, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		return
	}
	newTodo, db := models.GetTodoById(Id)
	if newTodo.Title == "" {
		newTodo.Title = updatetodo.Title
	}
	if newTodo.Name == "" {
		newTodo.Name = updatetodo.Name
	}
	if newTodo.Description == "" {
		newTodo.Description = updatetodo.Description
	}
	db.Save(&newTodo)
	res, _ := json.Marshal(newTodo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoId := params["id"]
	Id, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		return
	}
	todo := models.DeleteTodo(Id)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
