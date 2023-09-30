package mw

import (
	"github.com/gin-gonic/gin"
)

// Пример простенького middleware
// Пока что работает просто через header
// В дальнейшем надо будет сделать через токены
func IsAdminMW() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
