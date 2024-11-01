package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qa-app/controller"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to QA App")
	})
	userRouter := server.Group("/user")
	userRouter.POST("", userController.Create)
	userRouter.GET("", userController.FindAll)
	userRouter.POST("/login", userController.Login)
	return server
}
