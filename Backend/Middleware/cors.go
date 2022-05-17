package Middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.Default()
	// return cors.New(cors.Config{
	// 	AllowOrigins:  []string{"http://127.0.0.1"},
	// 	AllowMethods:  []string{"GET", "POST"},
	// 	AllowHeaders:  []string{"Origin"},
	// 	ExposeHeaders: []string{"Content-Length", "Authorization"},
	// AllowCredentials: true, 设置是否支持cookie跨域
	// 	MaxAge: 12 * time.Hour,
	// })
}
