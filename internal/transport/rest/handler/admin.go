package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary SignInAdmin
// @Tags auth
// @Description login for admins
// @ID login
// @Accept  json
// @Produce  json
// @Param input body SignInput true "admin info"
// @Success 200 {object} core.Tokens
// @Failure 500 {string} string "error"
// @Failure default {string} error "error"
// @Router /sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input SignInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.services.Admin.SignIn(input.Login, input.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// @Summary Refresh Access Token
// @Tags auth
// @Description Refresh access token for Admin
// @ID refresh
// @Accept  json
// @Produce  json
// @Param input body Refresh true "Refresh token"
// @Success 200 {string} string "token"
// @Failure 500 {string} string "error"
// @Failure default {string} error "error"
// @Router /refresh [post]
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
