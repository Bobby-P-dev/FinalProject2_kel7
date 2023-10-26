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
		photoRouter.GET("/getAll", controllers.GetPhoto)
		photoRouter.PUT("/edit/:id", middlewares.UserAuthorization(), controllers.EditPhoto)
		photoRouter.DELETE("/delete/:id", middlewares.UserAuthorization(), controllers.DeletePhoto)
	}
	return r
}
