package main

import (
	"github.com/Ibuki-Y/go-echo-clean/controller"
	"github.com/Ibuki-Y/go-echo-clean/db"
	"github.com/Ibuki-Y/go-echo-clean/repository"
	"github.com/Ibuki-Y/go-echo-clean/router"
	"github.com/Ibuki-Y/go-echo-clean/usecase"
	"github.com/Ibuki-Y/go-echo-clean/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
