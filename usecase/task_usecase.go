package usecase

import (
	"github.com/Ibuki-Y/go-echo-clean/model"
	"github.com/Ibuki-Y/go-echo-clean/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userID uint) ([]model.TaskResponse, error)
	GetTaskByID(userID, taskID uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userID, taskID uint) (model.TaskResponse, error)
	DeleteTask(userID, taskID uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUsecase{tr}
}

func (t *taskUsecase) GetAllTasks(userID uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := t.tr.GetAllTasks(&tasks, userID); err != nil {
		return nil, err
	}

	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}

	return resTasks, nil
}

func (t *taskUsecase) GetTaskByID(userID, taskID uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := t.tr.GetTaskByID(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (t *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := t.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (t *taskUsecase) UpdateTask(task model.Task, userID, taskID uint) (model.TaskResponse, error) {
	if err := t.tr.UpdateTask(&task, userID, taskID); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (t *taskUsecase) DeleteTask(userID, taskID uint) error {
	if err := t.tr.DeleteTask(userID, taskID); err != nil {
		return err
	}

	return nil
}
