package main

import (
	todo "learn/todoapi/pkg"
	"learn/todoapi/pkg/handler"
	"learn/todoapi/pkg/models"
	"learn/todoapi/pkg/service"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	db_config := viper.GetStringMapString("db")
	db, err := models.Connect(
		db_config["host"], db_config["port"], db_config["user"], 
		os.Getenv("DB_PASS"), db_config["dbname"],
	)
	if err != nil {
		logrus.Fatal(err)
	}

	repo := models.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	server := todo.NewServer(viper.GetString("port"), handler.InitRoutes())

	if err := server.Run(); err != nil {
		logrus.Fatal("error ocured while running server: ", err.Error())
	}
}
