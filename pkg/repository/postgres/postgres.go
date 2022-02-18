package postgres

import (
	"fmt"
	"learn/todoapi/pkg/models"
	"learn/todoapi/pkg/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	repository.Authorization
	repository.Task
	repository.TodoList

	db *gorm.DB
}

type Config struct {
	Host 		string 
	Port 		string
	User 		string 
	Password 	string
	DBName 		string
	SSLMode		string
}



func Connect(cfg Config) (repo *PostgresRepo, err error) {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode,
	)
	logrus.Debugf("connecting to postgres on %v:%v as %v to %v", 
				   cfg.Host, cfg.Port, cfg.User, cfg.DBName)

	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	repo = &PostgresRepo{
		Authorization: NewPgAuth(db),
		Task: NewPgTask(db),
		TodoList: NewPgList(db),
		db: db,
	}

	if err := db.AutoMigrate(userDB{}); err != nil {
		logrus.Fatal(err)
	}
	if err := db.AutoMigrate(models.Task{}); err != nil {
		logrus.Fatal(err)
	}
	if err := db.AutoMigrate(models.TodoList{}); err != nil {
		logrus.Fatal(err)
	}

	return
}

func (r *PostgresRepo) Close() error {
	sqldb, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqldb.Close()
}

