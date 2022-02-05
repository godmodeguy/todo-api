package main

import (
	todo "learn/todoapi/pkg"
	"learn/todoapi/pkg/handler"
	"learn/todoapi/pkg/models"
	"learn/todoapi/pkg/service"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error loading config")
	}

	repo := models.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	server := todo.NewServer(viper.GetString("port"), handler.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatal("error ocured while running server: ", err.Error())
	}
}
