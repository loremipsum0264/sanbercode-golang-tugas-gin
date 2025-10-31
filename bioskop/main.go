package main

import (
	"bioskop/controllers"
	"bioskop/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	err := database.Initiator()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := gin.Default()

    router.GET("/bioskop", controllers.GetAllBioskop)
    router.GET("/bioskop/:id", controllers.GetBioskop)
    router.POST("/bioskop", controllers.InsertBioskop)
    router.PUT("/bioskop/:id", controllers.UpdateBioskop)
    router.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}




