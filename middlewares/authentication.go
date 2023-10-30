package middlewares

import (
	"net/http"
	"strconv"

	"github.com/Bobby-P-dev/FinalProject2_kel7/database"
	"github.com/Bobby-P-dev/FinalProject2_kel7/helpers"
	"github.com/Bobby-P-dev/FinalProject2_kel7/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifToken, err := helpers.VerifToken(c)
		_ = verifToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unautenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifToken)
		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("users_id").First(&Photo, photoId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data Doesn`t Exist",
			})
			return
		}

		if Photo.UsersID != int(userID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to acces this data photo",
			})
			return
		}
		c.Next()
	}

}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("users_id").First(&Comment, commentId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data Doesn`t Exist",
			})
			return
		}

		if Comment.UsersID != int(userID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to acces this data photo",
			})
			return
		}
		c.Next()
	}

}

func SocialMAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		socialId, err := strconv.Atoi(c.Param("socialId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Social := models.SocialMedia{}

		err = db.Select("users_id").First(&Social, socialId).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data Doesn`t Exist",
			})
			return
		}

		if Social.UsersID != int(userID) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to acces this data photo",
			})
			return
		}
		c.Next()
	}

}
