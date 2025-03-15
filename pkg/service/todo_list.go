package service

import (
	todo "github.com/Xu3is/RestApiGarbage"
	"github.com/Xu3is/RestApiGarbage/pkg/repository"
)

type ToDoListService struct {
	repo repository.ToDoList
}

func NewToDoListService(repo repository.ToDoList) *ToDoListService {
	return &ToDoListService{repo: repo}
}
func (s *ToDoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ToDoListService) GetAll(userid int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userid)
}

func (s *ToDoListService) GetById(userId int, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *ToDoListService) Delete(userId int, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *ToDoListService) Update(userId int, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
