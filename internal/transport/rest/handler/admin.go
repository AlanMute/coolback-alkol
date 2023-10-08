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

	tokens, err := h.services.Admin.SignIn(input.Login, input.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) refreshAccess(c *gin.Context) {
	var refresh Refresh

	if err := c.ShouldBindJSON(&refresh); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.Refresh(refresh.RefreshToken)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
