package main

import (
	todo "learn/todoapi/pkg"
	"learn/todoapi/pkg/handler"
	"learn/todoapi/pkg/models"
	"learn/todoapi/pkg/service"
	"log"
)

func main() {

	repo := models.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	server := todo.NewServer("8090", handler.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatal("error ocured while running server: ", err.Error())
	}
}
