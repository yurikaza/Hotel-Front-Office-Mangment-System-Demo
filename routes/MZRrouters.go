package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/controllers"
)

//UserRoutes function
func MZRroutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("/MZR/verify", controller.MZRverify())
}