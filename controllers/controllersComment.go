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

func UploadComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	typeContent := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if typeContent == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UsersID = int(userID)

	err := db.Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"users_id":   Comment.UsersID,
		"created_at": Comment.CreatedAt,
	})
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	var Comment []models.Comment

	Comments := models.Comment{}
	userID := uint(userData["id"].(float64))

	Comments.UsersID = int(userID)
	err := db.Preload("Users").Preload("Photo").Find(&Comment, Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error",
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": Comment,
	})
}

func EditComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	typeContent := helpers.GetContentType(c)

	Comments := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if typeContent == appJSON {
		c.ShouldBindJSON(&Comments)
	} else {
		c.ShouldBind(&Comments)
	}

	Comments.UsersID = int(userID)
	Comments.ID = uint(commentId)

	err := db.Model(&Comments).Where("id = ?", commentId).Updates(models.Comment{
		Message: Comments.Message,
	}).Preload("Photo").First(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        Comments.ID,
		"message":   Comments.Message,
		"caption":   Comments.Photo.Caption,
		"photo_url": Comments.Photo.PhotoUrl,
		"title":     Comments.Photo.Title,
		"user_id":   Comments.UsersID,
		"update_at": Comments.UpdatedAt,
	})

}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comments := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	Comments.UsersID = int(userID)
	Comments.ID = uint(commentId)

	err := db.Model(&Comments).Where("id = ?", commentId).Delete(&Comments).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your Comment has been succesfully deleted",
	})
}
