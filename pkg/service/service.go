package service

import (
	"learn/todoapi/pkg/models"
)

type Authorization interface {
}

type TodoList interface {
}

type Task interface {
}

type Service struct {
	Authorization
	TodoList
	Task
}

func NewService(repo *models.Repository) *Service {
	return &Service{}
}
