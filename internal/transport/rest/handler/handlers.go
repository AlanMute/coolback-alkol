package handler

import (
	//"net/http"

	//"github.com/KrizzMU/coolback-alkol/repository"
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

	// GET запрос для swagger (later)

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
		course.Handle("POST", "/", h.AddCourse)
	}

	//router.Handle("GET", "/", handlers.TestHandler(repo))
	//router.Handle("POST", "/add/course", handlers.AddCourseHandler(repo))
	//router.Handle("POST", "/add/module", handlers.AddModuleHandler(repo))
	//router.Handle("POST", "/add/lesson", handlers.AddLessonHandler(repo))

	return r
}
