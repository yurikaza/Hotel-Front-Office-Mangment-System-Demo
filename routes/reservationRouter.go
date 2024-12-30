package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/controllers"
)

//UserRoutes function
func ReservationRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("/reservation/create", controller.CreateReservation())
}