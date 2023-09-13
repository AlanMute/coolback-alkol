package handlers

import (
	"net/http"

	"github.com/KrizzMU/coolback-alkol/repository"
	"github.com/gin-gonic/gin"
)

func TestHandler(repo repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	}
}
