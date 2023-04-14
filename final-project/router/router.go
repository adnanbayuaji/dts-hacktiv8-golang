package router

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)

		userRouter.POST("/", controllers.CreateUser)
		userRouter.PUT("/:userID", controllers.UpdateUser)
		userRouter.GET("/:userID", controllers.GetUser)
		userRouter.DELETE("/:userID", controllers.DeleteUser)
		userRouter.GET("/", controllers.GetUsers)
	}

	socialmediaRouter := r.Group("/socialmedias")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.GET("/", controllers.GetAllSocialmedia)
		socialmediaRouter.GET("/:socialmediaID", controllers.GetOneSocialmedia)
		socialmediaRouter.POST("/", controllers.CreateSocialmedia)
		socialmediaRouter.PUT("/:socialmediaID", middlewares.SocialmediaAuthorization(), controllers.UpdateSocialmedia)
		socialmediaRouter.DELETE("/:socialmediaID", middlewares.SocialmediaAuthorization(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:socialmediaID", controllers.GetOnePhoto)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.PUT("/:socialmediaID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:socialmediaID", middlewares.PhotoAuthorization(), controllers.DeleteUser)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.GET("/:socialmediaID", controllers.GetOneComment)
		commentRouter.POST("/:photoID", controllers.CreateComment)
		commentRouter.PUT("/:socialmediaID/:photoID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:socialmediaID", middlewares.CommentAuthorization(), controllers.DeleteUser)
	}

	return r
}
