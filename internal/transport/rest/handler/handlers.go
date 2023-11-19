package handler

import (
	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/KrizzMU/coolback-alkol/docs"
)

type Handler struct {
	tokenManger auth.TokenManager
	services    *service.Service
}

func NewHandler(s *service.Service, t auth.TokenManager) *Handler {
	return &Handler{
		services:    s,
		tokenManger: t,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Handle("POST", "/sign-in", h.signIn)

	r.Handle("POST", "/refresh", h.refreshAccess)

	adm := r.Group("/adm", h.isAdmin)
	{
		lesson := adm.Group("/lesson")
		{
			// Add swagger (client -> FormFile: file: file, name: string, description: string, courseName: string, moduleName: string. server -> json: {error: string})
			lesson.Handle("POST", "/", h.AddLesson)
			lesson.Handle("PUT", "/:id", h.EditLesson)
			lesson.Handle("DELETE", "/", h.DeleteLesson)
		}
		module := adm.Group("/module")
		{
			module.Handle("DELETE", "/", h.DeleteModule)
			// Add swagger (client -> json:{name: string, description: string, courseName: string}. server -> json: {error: string})
			module.Handle("POST", "/", h.AddModule)
			module.Handle("PUT", "/:id", h.EditModule)
			module.Handle("POST", "/image/:id", h.UploadModuleImage)
			module.Handle("DELETE", "/image/:id", h.DeleteModuleImage)
		}
		course := adm.Group("/course")
		{
			course.Handle("PUT", "/:id", h.EditCourse)
			// Add swagger (client -> json:{name: string, description: string}. server -> json: {error: string})
			course.Handle("POST", "/", h.AddCourse)
			// Add swagger (client -> json:{name: string}. server -> json: {error: string})
			course.Handle("DELETE", "/", h.DeleteCourse)
			course.Handle("POST", "/image/:id", h.UploadCourseImage)
			course.Handle("DELETE", "/image/:id", h.DeleteCourseImage)
		}
	}

	lesson := r.Group("/lesson")
	{
		lesson.Handle("GET", "/:orderid/:moduleid", h.GetLesson)
		lesson.Handle("POST", "/trial", h.SendTrialLesson)
		//lesson.Handle("GET", "/:orderid/:moduleid")
	}

	module := r.Group("/module")
	{
		module.Handle("GET", "/:id", h.GetModule)
		module.Handle("GET", "/image/:id", h.DownloadModuleImage)
		//module.Handle("GET", "/:moduleid")
	}

	course := r.Group("/course")
	{
		course.Handle("GET", "/search/:name", h.GetCourseByName) // Add swagger
		course.Handle("GET", "/getall/", h.GetAllCourses)        // Add swagger
		course.Handle("GET", "/:id", h.GetCourse)
		course.Handle("GET", "/image/:id", h.DownloadCourseImage)
	}

	return r
}
