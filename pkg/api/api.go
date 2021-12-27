package api

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-be/pkg/model"
	"todo-be/pkg/service"
)
/*
API struct
 */
type TodoAPI struct {
	TodoService service.TodoService
}
/*
Init TodoAPI
 */
func NewTodoAPI(c service.TodoService) TodoAPI {
	return TodoAPI{TodoService: c}
}
/*
Returns all tasks
 */
func (p TodoAPI) GetAllTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("HTTP method: %s", r.Method)
		log.Printf("Endpoint: %s", r.URL)
		log.Printf("Request header: %s", r.Header)

		tasks, _ := p.TodoService.AllTasks()
		log.Printf("Respond Data: %s", tasks)
		JSON(w, http.StatusOK, tasks)
	}
}
/*
Creates new task
 */
func (p TodoAPI) CreateNewTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("HTTP method: %s", r.Method)
		log.Printf("Endpoint: %s", r.URL)
		log.Printf("Request header: %s", r.Header)

		var tempModel model.Task
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&tempModel); err != nil {
			Error(w, http.StatusBadRequest, err, err.Error())
			return
		}
		defer r.Body.Close()

		createdPost, err := p.TodoService.CreateNewTask((&tempModel))
		log.Printf("Respond Data: %s", createdPost)

		if err != nil {
			Error(w, http.StatusInternalServerError, err, err.Error())
			return
		}
		JSON(w, http.StatusOK, createdPost)
	}
}
