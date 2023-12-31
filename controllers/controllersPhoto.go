package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bobby-P-dev/FinalProject2_kel7/database"
	"github.com/Bobby-P-dev/FinalProject2_kel7/helpers"
	"github.com/Bobby-P-dev/FinalProject2_kel7/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UploadPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	typeContent := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if typeContent == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UsersID = int(userID)

	err := db.Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err"     :     "bad request",
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id"        :        Photo.ID,
		"tilte"     :     Photo.Title,
		"caption"   :   Photo.Caption,
		"photo_url" : Photo.PhotoUrl,
		"user_id"   :   Photo.UsersID,
		"update_at" : Photo.UpdatedAt,
	})
}

func GetPhoto(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	var Photos []models.Photo

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	Photo.UsersID = int(userID)

	err := database.DB.Preload("Users").Find(&Photos, Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "error",
			"msg"   :   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"photo": Photos,
	})

}

func EditPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	typeContent := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if typeContent == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UsersID = int(userID)
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error"    : "Bad Request",
			"message"  : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id"        :        Photo.ID,
		"tilte"     :     Photo.Title,
		"caption"   :   Photo.Caption,
		"photo_url" : Photo.PhotoUrl,
		"user_id"   :   Photo.UsersID,
		"update_at" : Photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	Photo.UsersID = int(userID)
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error"   :   "unauthorized",
			"message" : "invalid",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your data photo has been succesfully deleted",
	})
}
