package models

import "gorm.io/gorm"

type Authorization interface {
}

type TodoList interface {
}

type Task interface {
}

type Repository struct {
	Authorization
	TodoList
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}
