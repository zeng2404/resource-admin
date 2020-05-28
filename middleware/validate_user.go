package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-admin/utils/log"
	"strings"
)

func ValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.String(), "login") {
			log.Info("验证通过")
			c.Next()
		} else {
			log.Info("验证失败")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
