package repository

import "learn/todoapi/pkg/models"


type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
	GetUserById(id int) (models.User, error)
}

type TodoList interface {
	GetAll() ([]models.TodoList, error)
	GetById(id int) (models.TodoList, error)
	CreateList(list models.TodoList) (int, error)
	UpdateList(list models.TodoList) (models.TodoList, error)
	DeleteList(id int) error
}

type Task interface {
}

type Repository struct {
	Authorization
	TodoList
	Task
}

func NewRepository(auth Authorization, list TodoList, task Task) Repository {
	return Repository{
		Authorization: auth,
		TodoList: list,
		Task: task,
	}
}





