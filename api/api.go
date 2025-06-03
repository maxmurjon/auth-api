package api

import (
	"github.com/maxmurjon/auth-api/api/handler"
	"github.com/maxmurjon/auth-api/config"

	"github.com/gin-gonic/gin"
	_ "github.com/maxmurjon/auth-api/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth API
// @version 1.0
// @description This is an authentication API
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func SetUpAPI(r *gin.Engine, h handler.Handler, cfg config.Config) {
	r.Use(customCORSMiddleware())

	r.POST("/login", h.Login)
	r.POST("/register", h.Register)

	userGroup := r.Group("/users")
	userGroup.Use(h.AuthMiddleware()) // middleware ni qo‘llash
	{
		userGroup.POST("/", h.CreateUser)
		userGroup.PUT("/", h.UpdateUser)
		userGroup.GET("/", h.GetUsersList)
		userGroup.GET("/:id", h.GetUsersByIDHandler)
		userGroup.DELETE("/:id", h.DeleteUser)
	}

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")                                                                                                      // Barcha manbalarga ruxsat berish
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")                                                          // Ruxsat etilgan metodlar
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSF-TOKEN, Authorization, Cache-Control") // So'rov sarlavhalari

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
