package controllers

import (
	"bioskop/database"
	"bioskop/repository"
	"bioskop/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBioskop(c *gin.Context) {
    var (
       result gin.H
    )

    bioskop, err := repository.GetAllBioskop(database.DbConnection)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": bioskop,
       }
    }

    c.JSON(http.StatusOK, result)
}

func GetBioskop(c *gin.Context) {
	var bioskop structs.Bioskop
    var (
       result gin.H
    )

    bioskop, err := repository.GetBioskop(database.DbConnection, bioskop)

    if err != nil {
       result = gin.H{
          "result": err.Error(),
       }
    } else {
       result = gin.H{
          "result": bioskop,
       }
    }

    c.JSON(http.StatusOK, result)
}

func InsertBioskop(c *gin.Context) {
    var bioskop structs.Bioskop

    err := c.BindJSON(&bioskop)
    if err != nil {
       panic(err)
    }

    err = repository.InsertBioskop(database.DbConnection, bioskop)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, bioskop)
}

func UpdateBioskop(c *gin.Context) {
    var bioskop structs.Bioskop
    id, _ := strconv.Atoi(c.Param("id"))

    err := c.BindJSON(&bioskop)
    if err != nil {
       panic(err)
    }

    bioskop.ID = id

    err = repository.UpdateBioskop(database.DbConnection, bioskop)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, bioskop)
}

func DeleteBioskop(c *gin.Context) {
    var bioskop structs.Bioskop
    id, _ := strconv.Atoi(c.Param("id"))

    bioskop.ID = id
    err := repository.DeleteBioskop(database.DbConnection, bioskop)
    if err != nil {
       panic(err)
    }

    c.JSON(http.StatusOK, bioskop)
}