package main

import (
	"github.com/gin-gonic/gin"
	"openrdv-server/models"
	"openrdv-server/controllers"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/devices", controllers.FindDevices)

	err := r.Run()
	if err != nil {

	}
}