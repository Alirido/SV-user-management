package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Alirido/SV-user-management/models"
	"github.com/Alirido/SV-user-management/controllers"
)

func main() {
	route := gin.Default()

	models.ConnectDatabase()

	route.GET("/users", controllers.GetUsers)
	route.GET("/users/:id", controllers.GetUser)
	route.POST("/users", controllers.CreateUser)
	route.PATCH("/users/:id", controllers.UpdateUser)
	route.DELETE("/users/:id", controllers.DeleteUser)

	route.Run()
}
