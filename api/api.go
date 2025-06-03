package api

import (
	"smartlogistics/api/handler"
	"smartlogistics/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "smartlogistics/api/docs" 
)

func SetUpAPI(r *gin.Engine, h handler.Handler, cfg config.Config) {
	r.Use(customCORSMiddleware())

	// Users Endpoints
	r.POST("/createuser", h.CreateUser)
	r.PUT("/updateuser", h.UpdateUser)
	r.GET("/users", h.GetUsersList)
	r.GET("/user/:id", h.GetUsersByIDHandler)
	r.DELETE("/deleteuser/:id", h.DeleteUser)

	//Product Endpoints
	r.POST("/createproduct", h.CreateProduct)
	r.PUT("/updateproduct", h.UpdateProduct)
	r.GET("/products", h.GetProductsList)
	r.GET("/product/:id", h.GetProductsByIDHandler)
	r.DELETE("/deleteproduct/:id", h.DeleteProduct)

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
