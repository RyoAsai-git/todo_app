package main

import (
	"fmt"
	"github.com/RyoAsai-git/todo_app/models"
	// "github.com/RyoAsai-git/todo_app/config"
	// "log"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)
}
