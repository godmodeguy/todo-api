package service

import (
	"learn/todoapi/pkg/models"
	"learn/todoapi/pkg/repository"
)

type Authorization interface {
	CreateUser(models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(models.TodoList) (int, error)
	GetAll() ([]models.TodoList, error)
	GetById(int) (models.TodoList, error)
	UpdateList(list models.TodoList) (models.TodoList, error)
	DeleteList(int) error
}

type Task interface {
}

type Service struct {
	Authorization
	TodoList
	Task
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList: NewListService(repo.TodoList),
		Task: NewTaskService(repo.Task),
	}
}
