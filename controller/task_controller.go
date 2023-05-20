package controller

import (
	"net/http"
	"strconv"

	"github.com/Ibuki-Y/go-echo-clean/model"
	"github.com/Ibuki-Y/go-echo-clean/usecase"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskByID(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (t *taskController) GetAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]

	tasksRes, err := t.tu.GetAllTasks(uint(userID.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasksRes)
}

func (t *taskController) GetTaskByID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	id := c.Param("taskID")
	taskID, _ := strconv.Atoi(id)

	tasksRes, err := t.tu.GetTaskByID(uint(userID.(float64)), uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasksRes)
}

func (t *taskController) CreateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	task.UserID = uint(userID.(float64))
	taskRes, err := t.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, taskRes)
}

func (t *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	id := c.Param("taskID")
	taskID, _ := strconv.Atoi(id)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	taskRes, err := t.tu.UpdateTask(task, uint(userID.(float64)), uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (t *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"]
	id := c.Param("taskID")
	taskID, _ := strconv.Atoi(id)

	err := t.tu.DeleteTask(uint(userID.(float64)), uint(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
