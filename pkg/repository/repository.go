package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Repository struct {
	Authorization
	ToDoList
	ToDoItem
}

func NewRepository(s *sqlx.DB) *Repository {
	return &Repository{}
}
