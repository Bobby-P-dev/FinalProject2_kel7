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

func UploadSocialMedia(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	typeContent := helpers.GetContentType(c)

	Social := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if typeContent == appJSON {
		c.ShouldBindJSON(&Social)
	} else {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	Social.UsersID = int(userID)

	err := db.Create(&Social).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     err.Error(),
			"message": "Failed to create data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":               Social.ID,
		"name":             Social.Name,
		"social_media_url": Social.SocialMediaUrl,
		"users_id":         Social.UsersID,
		"created_at":       Social.CreatedAt,
	})
}

func GetSocialMedia(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var Social []models.SocialMedia
	SocialM := models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	SocialM.UsersID = int(userID)

	err := db.Preload("Users").Find(&Social, SocialM).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "Failed to get data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Social,
	})
}

func EditSocialMedia(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	typeContent := helpers.GetContentType(c)

	Social := models.SocialMedia{}

	soacialId, _ := strconv.Atoi(c.Param("socialId"))
	userID := uint(userData["id"].(float64))

	if typeContent == appJSON {
		c.ShouldBindJSON(&Social)
	} else {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	Social.UsersID = int(userID)
	Social.ID = uint(soacialId)

	err := db.Model(&Social).Where("id = ?", soacialId).Updates(models.SocialMedia{
		Name: Social.Name, SocialMediaUrl: Social.SocialMediaUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to update data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":               Social.ID,
		"name":             Social.Name,
		"social_media_url": Social.SocialMediaUrl,
		"user_id":          Social.UsersID,
		"update_at":        Social.UpdatedAt,
	})

}

func DeleteSocialMedia(c *gin.Context) {

	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Social := models.SocialMedia{}

	soacialId, _ := strconv.Atoi(c.Param("socialId"))
	userID := uint(userData["id"].(float64))

	Social.UsersID = int(userID)
	Social.ID = uint(soacialId)

	err := db.Model(&Social).Where("id = ?", soacialId).Delete(&Social).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Failed to delete data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your data Social Media has been succesfully deleted",
	})
}
