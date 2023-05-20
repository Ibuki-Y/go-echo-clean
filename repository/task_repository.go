package repository

import (
	"fmt"

	"github.com/Ibuki-Y/go-echo-clean/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userID uint) error
	GetTaskByID(task *model.Task, userID, taskID uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userID, taskID uint) error
	DeleteTask(userID, taskID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) GetAllTasks(tasks *[]model.Task, userID uint) error {
	if err := t.db.Joins("User").Where("user_id=?", userID).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetTaskByID(task *model.Task, userID, taskID uint) error {
	if err := t.db.Joins("User").Where("user_id=?", userID).First(task, taskID).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) CreateTask(task *model.Task) error {
	if err := t.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) UpdateTask(task *model.Task, userID, taskID uint) error {
	result := t.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskID, userID).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (t *taskRepository) DeleteTask(userID, taskID uint) error {
	result := t.db.Where("id=? AND user_id=?", taskID, userID).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
