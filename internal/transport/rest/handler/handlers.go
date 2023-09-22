package handler

import (
	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	lesson := r.Group("/lesson")
	{
		lesson.Handle("GET", "/:id", h.GetLesson)
		lesson.Handle("POST", "/", h.AddLesson)
	}

	module := r.Group("/module")
	{
		module.Handle("POST", "/", h.AddModule)
	}

	course := r.Group("/course")
	{
		course.Handle("POST", "/", h.AddCourse)           // Add swagger
		course.Handle("GET", "/:name", h.GetCourseByName) // Add swagger
		course.Handle("GET", "/", h.GetAllCourses)        // Add swagger
	}

	return r
}
