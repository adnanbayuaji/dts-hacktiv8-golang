package router

import (
	"challenge-8/controllers"
	"challenge-8/middlewares"

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

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
