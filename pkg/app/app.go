package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"learn/todoapi/pkg/handler"
	"learn/todoapi/pkg/repository"
	"learn/todoapi/pkg/repository/postgres"
	"learn/todoapi/pkg/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Start() {
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)


	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	if err := godotenv.Load("./configs/.env"); err != nil {
		logrus.Fatal(err)
	}
	if os.Getenv("JWT_SECRET") == "" || 
	   os.Getenv("HASH_SALT") == "" || 
	   os.Getenv("DB_POSTGRES_PASSWORD") == "" {
		logrus.Fatal("set up secrets in ./configs/.env")
	}


	db_config := viper.GetStringMapString("db")
	cfg := postgres.Config{
		Host: 		db_config["host"], 
		Port: 		db_config["port"], 
		User: 		db_config["user"], 
		Password: 	os.Getenv("DB_POSTGRES_PASSWORD"), 
		DBName: 	db_config["dbname"],
		SSLMode: 	db_config["sslmode"],
	}

	pg_repo, err := postgres.Connect(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository(pg_repo.Authorization, pg_repo.TodoList, pg_repo.Task)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	server := NewServer(viper.GetString("port"), handler.InitRoutes())

	go server.Run()

	logrus.Printf("Here we go! Server started. PID [%v]\n", os.Getpid())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	
	switch <-quit {
	case syscall.SIGINT:
		logrus.Println("Stopped by Ctrl-C. Bye!")
	case syscall.SIGTERM:
		logrus.Println("Are you killing me? Well, i'll die.")
	}

	logrus.Println("Shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorln("error while shutting down server: ", err.Error())
	}

	if err := pg_repo.Close(); err != nil {
		logrus.Errorln("error closing connection with DB: ", err.Error())
	}
}
