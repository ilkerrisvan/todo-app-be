package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"todo-be/pkg/api"
	"todo-be/pkg/repository"
	"todo-be/pkg/service"
)
/*
Todo App
 */
type TodoApp struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	app := TodoApp{}
	err = app.initialize(os.Getenv("CONNECTION_STRING"))
	if err != nil{
		log.Println("Todo APP can't initialize.")
	}
	app.routes()
	err = app.run(os.Getenv("SERV_PORT"))
}

func (app *TodoApp) initialize(connectionAdrr string) error {
	var err error
	connectionString := fmt.Sprintf("%s", connectionAdrr)
	app.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return err
	}
	app.Router = mux.NewRouter()
	return err
}

func (app *TodoApp) run(addr string) error {
	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		log.Fatalf("Port is using already.")
		return err
	}
	return err
}

func (app *TodoApp) routes() {
	todoAPI := InitConverterAPI(app.DB)
	app.Router.HandleFunc("/todo", todoAPI.GetAllTasks()).Methods("GET", "OPTIONS")
	app.Router.HandleFunc("/todo", todoAPI.CreateNewTask()).Methods("POST", "OPTIONS")
}

func InitConverterAPI(db *gorm.DB) api.TodoAPI {
	todoRepository := repository.NewRepository(db)
	todoService := service.NewTodoService(todoRepository)
	todoAPI := api.NewTodoAPI(todoService)
	return todoAPI
}

