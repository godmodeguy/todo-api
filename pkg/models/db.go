package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(
	host string, 
	port string,
	user string, 
	password string,
	dbName string,
) (DB *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		host, user, password, dbName, port,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return
}