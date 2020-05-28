package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-admin/utils/log"
)

// 处理跨域 middleware
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		log.Infof("the method is %v\n", method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// OPTIONS 直接放行,因为有的模板会请求两次
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 进一步处理请求
		c.Next()
	}
}
