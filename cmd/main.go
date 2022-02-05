package main

import (
	todo "learn/todoapi/pkg"
	"log"
)

func main() {
	server := todo.NewServer(8090)
	if err := server.Run(); err != nil {
		log.Fatal("error ocured while running server: ", err.Error())
	}
}