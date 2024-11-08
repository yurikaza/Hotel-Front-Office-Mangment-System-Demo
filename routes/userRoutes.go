package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/controllers"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("/users/signup", controller.SignUp())
    incomingRoutes.POST("/users/login", controller.Login())
}