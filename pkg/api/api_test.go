package api

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"todo-be/pkg/model"
	"todo-be/pkg/repository"
	"todo-be/pkg/service"

	"github.com/stretchr/testify/assert"
)

func dbSetup() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	DB, _ := gorm.Open("postgres", db)
	DB.LogMode(true)
	return DB, mock
}

func routersSetup(api TodoAPI) *mux.Router {
	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/todo", api.GetAllTasks()).Methods("GET")
	apiRouter.HandleFunc("/todo", api.GetAllTasks()).Methods("POST")
	return apiRouter
}

func apiSetup(db *gorm.DB) TodoAPI {
	postRepository := repository.NewRepository(db)
	postService := service.NewTodoService(postRepository)
	todoAPI := NewTodoAPI(postService)
	return todoAPI
}

func TestGetAllTasks(t *testing.T) {
	w := httptest.NewRecorder()
	mockDB, mock := dbSetup()
	api := apiSetup(mockDB)
	r := routersSetup(api)

	var tasks []model.Task
	task := model.Task{
		Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}

	tasks = append(tasks, task)

	rows := sqlmock.
		NewRows([]string{"text"}).
		AddRow(task.Text)
	const sqlSelectOne = `SELECT * FROM "tasks"`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).
		WillReturnRows(rows)

	r.ServeHTTP(w, httptest.NewRequest("GET", "/todo", nil))
	assert.Equal(t, http.StatusOK, w.Code, "Did not get expected HTTP status code, got")
	var resultPosts []model.Task
	decoder := json.NewDecoder(w.Body)
	if err := decoder.Decode(&resultPosts); err != nil {
		t.Error(err)
	}
	resultPosts[0].Text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	assert.Equal(t, tasks, resultPosts)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
func TestCreateNewTask(t *testing.T) {
	w := httptest.NewRecorder()
	mockDB, mock := dbSetup()
	api := apiSetup(mockDB)
	r := routersSetup(api)

	var tasks []model.Task
	task := model.Task{
		Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}

	tasks = append(tasks, task)
	rows := sqlmock.
		NewRows([]string{"text"}).
		AddRow(task.Text)
	const sqlSelectOne = `INSERT INTO tasks(text) VALUES (Lorem ipsum dolor sit amet, consectetur adipiscing elit.)`
	mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).
		WillReturnRows(rows)

	r.ServeHTTP(w, httptest.NewRequest("POST", "/todo", nil))

}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
