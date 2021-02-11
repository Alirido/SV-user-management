package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Alirido/SV-user-management/models"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateUserInput struct {
  Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`  
}

// Endpoint: GET /users
// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Endpoint: GET /users/:id
// Get an user
func GetUser(c *gin.Context) {
  var user models.User

  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}

// Endpoint: POST /users
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Username: input.Username, Password: input.Password, Name: input.Name}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Endpoint: PATCH /users/:id
// Update an user
func UpdateUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input UpdateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
	}

	// Update and Save user
	models.DB.Model(&user).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": user})
}