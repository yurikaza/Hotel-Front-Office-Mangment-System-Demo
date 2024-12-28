package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yurikaza/Hotel-Front-Office-Mangment-System-Demo-Backend/routes"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        port = "8001"
    }

    router := gin.New()

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "https://github.com"
        },
        MaxAge: 12 * time.Hour,
    }))

    router.Use(gin.Logger())

    routes.MZRroutes(router)

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