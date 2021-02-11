package controllers

import (
	"net/http"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/Alirido/SV-user-management/models"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=7"`
	Name string `json:"name" binding:"required,min=3"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name""`  
}

// BeforeUpdate : hook before a user is updated
func (u *UpdateUserInput) BeforeUpdate(scope *gorm.Scope) (err error) {
	fmt.Println("before update")
	fmt.Println(u.Password)

	if u.Password != "" && len(u.Password) >= 7 {
			hash, err := HashAndSaltPwd(u.Password)
			if err != nil {
					return nil
			}
			scope.SetColumn("Password", hash)
	}

	fmt.Println(u.Password)
	return
}

func HashAndSaltPwd(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashedPassword), err
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

	hashedPassword, err := HashAndSaltPwd(input.Password)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
	}

	// Create user
	user := models.User{Username: input.Username, Password: hashedPassword, Name: input.Name}
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

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&user)

  c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted user with ID:" + c.Param("id")})
}