package main

import (
	"github.com/Nikita-Ignatenkov/todo-app"
	"github.com/Nikita-Ignatenkov/todo-app/pkg/handler"
	"log"
)

func main() {
	hendlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", hendlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}
