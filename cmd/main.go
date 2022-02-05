package main

import (
	todo "learn/todoapi/pkg"
	"learn/todoapi/pkg/handler"
	"log"
)

func main() {
	handler := handler.NewHandler()

	server := todo.NewServer(8090, handler.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatal("error ocured while running server: ", err.Error())
	}
}