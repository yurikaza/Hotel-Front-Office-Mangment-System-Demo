package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/middleware"
	"github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/routes"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())
    routes.UserRoutes(router)

    router.Use(middleware.Authentication())

    // API-2
    router.GET("/api-1", func(c *gin.Context) {

        c.JSON(200, gin.H{"success": "Access granted for api-1"})

    })

    // API-1
    router.GET("/api-2", func(c *gin.Context) {
        c.JSON(200, gin.H{"success": "Access granted for api-2"})
    })

    router.Run(":" + port)
}