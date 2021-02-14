package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Alirido/SV-user-management/models"
	"github.com/Alirido/SV-user-management/controllers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

			if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
			}

			c.Next()
	}
}

func main() {
	route := gin.Default()
	route.Use(CORSMiddleware())

	models.ConnectDatabase()

	route.GET("/users", controllers.GetUsers)
	route.GET("/users/:id", controllers.GetUser)
	route.POST("/users", controllers.CreateUser)
	route.PATCH("/users/:id", controllers.UpdateUser)
	route.DELETE("/users/:id", controllers.DeleteUser)

	route.Run()
}
