package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userRole(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Empty auth header!")
		return
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid auth header!")
		return
	}

	if len(parts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Tocken is empty!")
		return
	}

}
