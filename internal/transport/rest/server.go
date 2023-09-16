package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, ginEngine *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        ginEngine,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
