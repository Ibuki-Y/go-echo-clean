package migrate

import (
	"fmt"

	"github.com/Ibuki-Y/go-echo-clean/db"
	"github.com/Ibuki-Y/go-echo-clean/model"
)

func Migrate() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.Close(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
