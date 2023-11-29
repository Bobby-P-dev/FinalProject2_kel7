package controllers

import (
	"net/http"

	"github.com/Bobby-P-dev/FinalProject2_kel7/database"
	"github.com/Bobby-P-dev/FinalProject2_kel7/helpers"
	"github.com/Bobby-P-dev/FinalProject2_kel7/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	appJSON = "application/json"
)

func RegisterUser(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.Users{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	err := db.Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid to created account",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})

}

func LoginUser(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	User := models.Users{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	password = User.Password

	err := db.Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "invalid email",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "invalid password",
		})
		return
	}
	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func EditUser(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	id := c.Param("id")

	User := models.Users{}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	User.ID = userID

	err := db.Model(&User).Where("id = ?", id).Updates(models.Users{
		Email: User.Email, Username: User.Username,
	}).First(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  err.Error(),
			"message": "failed to edit user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"age":       User.Age,
		"id":        User.ID,
		"email":     User.Email,
		"username":  User.Username,
		"update_at": User.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	User := models.Users{}

	userID := uint(userData["id"].(float64))

	User.ID = userID

	err := db.Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "invalid Delete Data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been  succesfully deleted",
	})
}
