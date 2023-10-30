package router

import (
	"github.com/Bobby-P-dev/FinalProject2_kel7/controllers"
	"github.com/Bobby-P-dev/FinalProject2_kel7/middlewares"
	"github.com/gin-gonic/gin"
)

func StarApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
		userRouter.PUT("/edit/:id", middlewares.Authentication(), controllers.EditUser)
		userRouter.DELETE("/delete", middlewares.Authentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/upload", controllers.UploadPhoto)
		photoRouter.GET("/get", controllers.GetPhoto)
		photoRouter.PUT("/put/:photoId", middlewares.UserAuthorization(), controllers.EditPhoto)
		photoRouter.DELETE("/del/:photoId", middlewares.UserAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/upload", controllers.UploadComment)
		commentRouter.GET("/get", controllers.GetComment)
		commentRouter.PUT("/put/:commentId", middlewares.CommentAuthorization(), controllers.EditComment)
		commentRouter.DELETE("/del/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialRouter := r.Group("/socialm")
	{
		socialRouter.Use(middlewares.Authentication())
		socialRouter.POST("/upload", controllers.UploadSocialMedia)
		socialRouter.GET("/get", controllers.GetSocialMedia)
		socialRouter.PUT("/put/:socialId", middlewares.SocialMAuthorization(), controllers.EditSocialMedia)
		socialRouter.DELETE("/del/:socialId", middlewares.SocialMAuthorization(), controllers.DeleteSocialMedia)
	}
	return r
}
