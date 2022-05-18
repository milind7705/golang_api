package routes

import (
	"golang_api/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", controller.GetUsers)
		grp1.POST("user", controller.CreateUser)
		grp1.GET("user/:id", controller.GetUserByID)
		grp1.PUT("user/:id", controller.UpdateUser)
		grp1.DELETE("user/:id", controller.DeleteUser)
	}
	return r
}
