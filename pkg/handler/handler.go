package handler

import (
	"github.com/Xu3is/RestApiGarbage/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		courses := api.Group("/courses")
		{
			courses.GET("/courses", h.getCourses)
			courses.POST("/courses", h.createCourse)
			courses.PUT("/courses/:id", h.updateCourse)
			courses.DELETE("/courses/:id", h.deleteCourse)
		}
		trainers := api.Group("/trainers")
		{
			trainers.GET("/trainers", h.getTrainers)
			trainers.POST("/trainers", h.createTrainer)
			trainers.PUT("/trainers/:id", h.updateTrainer)
			trainers.DELETE("/trainers/:id", h.deleteTrainer)
		}
		lessons := api.Group("/lessons")
		{
			lessons.GET("/lessons", h.getLessons)
			lessons.POST("/lessons", h.createLesson)
			lessons.PUT("/lessons/:id", h.updateLesson)
			lessons.DELETE("/lessons/:id", h.deleteLesson)
		}
		students := api.Group("/students")
		{
			students.GET("/students", h.getStudents)
			students.POST("/students", h.createStudent)
			students.PUT("/students/:id", h.updateStudent)
			students.DELETE("/students/:id", h.deleteStudent)
		}
		payments := api.Group("/payments")
		{
			payments.GET("/payments", h.getPayments)
			payments.POST("/payments", h.createPayment)
			payments.PUT("/payments/:id", h.updatePayment)
			payments.DELETE("/payments/:id", h.deletePayment)
		}
	}

	return router
}
