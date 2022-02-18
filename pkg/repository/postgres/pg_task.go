package postgres

import (
	"gorm.io/gorm"
)


type PgTask struct {
	db	*gorm.DB
}

func NewPgTask(db *gorm.DB) PgTask {
	return PgTask{db: db}
}
