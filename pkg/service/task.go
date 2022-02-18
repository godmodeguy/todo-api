package service

import "learn/todoapi/pkg/repository"

type TaskService struct {}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{}
}