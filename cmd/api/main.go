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
	userUsecase := usecase.NewUserUsecase(userRepository)
	useController := controller.NewUserController(userUsecase)
	e := router.NewRouter(useController)
	e.Logger.Fatal(e.Start(":8080"))
}
