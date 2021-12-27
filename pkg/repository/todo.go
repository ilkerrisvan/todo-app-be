package repository

import (
	"github.com/jinzhu/gorm"
	"todo-be/pkg/model"
)
type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AllTasks() ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (p *Repository) CreateNewTask(post *model.Task) (*model.Task, error) {
	err := p.db.Save(&post).Error
	return post, err
}