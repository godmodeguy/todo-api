package service

import (
	"learn/todoapi/pkg/models"
	"learn/todoapi/pkg/repository"
)

type ListService struct {
	repo repository.TodoList
}

func NewListService(repo repository.TodoList) *ListService {
	return &ListService{repo: repo}
}

func (ls *ListService) CreateList(list models.TodoList) (int, error) {
	return ls.repo.CreateList(list)
}

func (ls *ListService) GetAll() ([]models.TodoList, error) {
	return ls.repo.GetAll()
}

func (ls *ListService) GetById(id int) (models.TodoList, error) {
	return ls.repo.GetById(id)
}

func (ls *ListService) UpdateList(list models.TodoList) (models.TodoList, error) {
	return ls.repo.UpdateList(list)
}

func (ls *ListService) DeleteList(id int) error {
	return ls.repo.DeleteList(id)
}