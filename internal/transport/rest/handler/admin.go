package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn(c *gin.Context) {
	var input SignInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
