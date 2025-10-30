package main

import (
	"bioskop/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bioskop struct {
	ID     int     `json:"id"`
	Nama   string  `json:"nama" binding:"required"`
	Lokasi string  `json:"lokasi" binding:"required"`
	Rating float64 `json:"rating"`
}

func main() {

	err := database.Initiator()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := gin.Default()

	router.POST("/bioskop", PostBioskop)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}

func PostBioskop(c *gin.Context) {
	var bioskop Bioskop

	// Bind JSON input ke struct
	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
			"details": err.Error(),
		})
		return
	}

	// Validasi Nama dan Lokasi
	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama and Lokasi dibutuhkan",
		})
		return
	}

	// Insert ke database
	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := database.DBConnections.QueryRow(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal membuat data bioskop",
		})
		return
	}

	bioskop.ID = id
	c.JSON(http.StatusCreated, gin.H{
		"message": "Data bioskop berhasil dibuat",
		"data":    bioskop,
	})
}


