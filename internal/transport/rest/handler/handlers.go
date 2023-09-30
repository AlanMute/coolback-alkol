package handler

import (
	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/KrizzMU/coolback-alkol/internal/transport/rest/mw"
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
		// Add swagger (client -> FormFile: file: file, name: string, description: string, courseName: string, moduleName: string. server -> json: {error: string})
		lesson.Handle("POST", "/", mw.IsAdminMW(), h.AddLesson)

		lesson.Handle("DELETE", "/", mw.IsAdminMW(), h.DeleteLesson)
	}

	module := r.Group("/module")
	{
		// Add swagger (client -> json:{name: string, description: string, courseName: string}. server -> json: {error: string})
		module.Handle("POST", "/", mw.IsAdminMW(), h.AddModule)
		module.Handle("DELETE", "/", mw.IsAdminMW(), h.DeleteModule)
		//module.Handle("GET", "/:moduleid")
	}

	course := r.Group("/course")
	{
		// Add swagger (client -> json:{name: string, description: string}. server -> json: {error: string})
		course.Handle("POST", "/", mw.IsAdminMW(), h.AddCourse)
		// Add swagger (client -> json:{name: string}. server -> json: {error: string})
		course.Handle("DELETE", "/", mw.IsAdminMW(), h.DeleteCourse)
		course.Handle("GET", "/:name", h.GetCourseByName) // Add swagger
		course.Handle("GET", "/", h.GetAllCourses)        // Add swagger
		//course.Handle("GET", "/:courseid")
	}
	//    local/Golang-for-begginer/Begin/Lets-start
	r.Handle("GET", "/:coursename/:modulename/:lessonname", h.GetLesson)

	r.Handle("GET", "/:coursename/:modulename/", h.GetModule)

	r.Handle("GET", "/:coursename/", h.GetCourse)
	return r
}
