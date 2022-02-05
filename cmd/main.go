package main

import (
	todo "learn/todoapi/pkg"
	"learn/todoapi/pkg/handler"
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	port := 6666

	if envs, err := godotenv.Read(".env"); err == nil {
		log.Println(".env file loaded")

		if p, err := strconv.ParseInt(envs["PORT"], 10, 32); err == nil {
			port = int(p)
			log.Println("port set to ", port)
		} else {
			log.Fatal("error loading port from .env")
		}

	} else {
		log.Fatal("error loading .env file")
	}

	handler := handler.NewHandler()

	server := todo.NewServer((port), handler.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatal("error ocured while running server: ", err.Error())
	}
}