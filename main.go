package main

import (
	"fmt"
	"log"
	"todo_app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DBName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
