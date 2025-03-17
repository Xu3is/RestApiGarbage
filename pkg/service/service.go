package service

import "github.com/Xu3is/RestApiGarbage/pkg/repository"

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

type Service struct {
	Authorization Authorization
	Courses       Courses
	Lessons       Lessons
	Payments      Payments
	Students      Students
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
