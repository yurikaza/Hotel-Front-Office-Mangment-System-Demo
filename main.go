package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/routes"
)

func main() {
    port := os.Getenv("PORT")

    if port == "8000" {
        port = "8001"
    }

    router := gin.New()

   	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

    router.Use(gin.Logger())

    routes.MZRroutes(router)
    routes.ReservationRoutes(router)

    /* 
    router.Use(middleware.Authentication())

    // API-2
    router.GET("/api-1", func(c *gin.Context) {

        c.JSON(200, gin.H{"success": "Access granted for api-1"})

    })

    // API-1
    router.GET("/api-2", func(c *gin.Context) {
        c.JSON(200, gin.H{"success": "Access granted for api-2"})
    })
    */

    router.Run(":" + port)
}