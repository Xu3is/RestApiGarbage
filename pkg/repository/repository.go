package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}
type Courses interface {
}
type Lessons interface {
}
type Payments interface {
}
type Trainers interface {
}

type Students interface {
}

type Repository struct {
	Authorization Authorization
	Courses       Courses
	Lessons       Lessons
	Payments      Payments
	Students      Students
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
