package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ritik-agarwal50/go_api_tutorial/pkg/config"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Title       string `json:"title"`   //1
	Name        string `json:"name"`  //3
	Description string `json:"description"` //2
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func (b *Todo) CreateTodo() *Todo {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTodos() []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos
}

func GetTodoById(Id int64) (*Todo, *gorm.DB) {
	var getTodo Todo
	db := db.Where("ID = ?", Id).Find(&getTodo)
	return &getTodo, db
}
func DeleteTodo(Id int64) Todo {
	var todo Todo
	db.Where("ID = ?", Id).Delete(todo)
	return todo
}
