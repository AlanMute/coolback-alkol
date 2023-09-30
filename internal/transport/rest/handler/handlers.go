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
		lesson.Handle("GET", "/:orderid/:moduleid", h.GetLesson)
		// Add swagger (client -> FormFile: file: file, name: string, description: string, courseName: string, moduleName: string. server -> json: {error: string})
		lesson.Handle("POST", "/", h.AddLesson)
		//lesson.Handle("GET", "/:orderid/:moduleid")
		lesson.Handle("DELETE", "/", h.DeleteLesson)
	}

	module := r.Group("/module")
	{
		// Add swagger (client -> json:{name: string, description: string, courseName: string}. server -> json: {error: string})
		module.Handle("POST", "/", h.AddModule)
		module.Handle("DELETE", "/", h.DeleteModule)
		module.Handle("GET", "/:id", h.GetModule)
		//module.Handle("GET", "/:moduleid")
	}

	course := r.Group("/course")
	{
		// Add swagger (client -> json:{name: string, description: string}. server -> json: {error: string})
		course.Handle("POST", "/", h.AddCourse)
		// Add swagger (client -> json:{name: string}. server -> json: {error: string})
		course.Handle("DELETE", "/", h.DeleteCourse)
		course.Handle("GET", "/search/:name", h.GetCourseByName) // Add swagger
		course.Handle("GET", "/search/", h.GetAllCourses)        // Add swagger
		course.Handle("GET", "/:id", h.GetCourse)
	}

	return r
}
