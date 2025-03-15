package service

import (
	todo "github.com/Xu3is/RestApiGarbage"
	"github.com/Xu3is/RestApiGarbage/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	Create(UserId int, list todo.TodoList) (int, error)
	GetAll(userid int) ([]todo.TodoList, error)
	GetById(userid int, listId int) (todo.TodoList, error)
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ToDoList:      NewToDoListService(repos.ToDoList),
	}
}
