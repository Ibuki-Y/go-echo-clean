package main

import (
	"github.com/Ibuki-Y/go-echo-clean/controller"
	"github.com/Ibuki-Y/go-echo-clean/db"
	"github.com/Ibuki-Y/go-echo-clean/repository"
	"github.com/Ibuki-Y/go-echo-clean/router"
	"github.com/Ibuki-Y/go-echo-clean/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
