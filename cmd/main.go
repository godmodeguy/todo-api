package main

import (
	"github.com/joho/godotenv"
	todo "learn/todoapi/pkg"
	"learn/todoapi/pkg/handler"
	"log"
	"strconv"
)

func main() {
	var port int64
	var envs map[string]string
	var err error

	if envs, err = godotenv.Read(".env"); err != nil {
		log.Fatal("error loading .env file")
	}
	log.Println(".env file loaded")

	if port, err = strconv.ParseInt(envs["PORT"], 10, 32); err != nil {
		log.Fatal("error loading port from .env")
	}
	log.Println("port set to ", port)

	handler := handler.NewHandler()

	server := todo.NewServer(int(port), handler.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatal("error ocured while running server: ", err.Error())
	}
}
