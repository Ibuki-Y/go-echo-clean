package router

import (
	"os"

	"github.com/Ibuki-Y/go-echo-clean/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskID", tc.GetTaskByID)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskID", tc.UpdateTask)
	t.DELETE("/:taskID", tc.DeleteTask)

	return e
}
