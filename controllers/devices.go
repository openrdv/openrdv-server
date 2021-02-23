package controllers

import (
	"github.com/gin-gonic/gin"
	"openrdv-server/models"
	"net/http"
)

type CreateDeviceInput struct {

}

func FindDevices(c *gin.Context) {
	var devices []models.Device
	models.DB.Find(&devices)

	c.JSON(http.StatusOK, gin.H{"data": devices})
}

func CreateDevice(c *gin.Context) {
	var input CreateDeviceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create device
	device := models.Device{}
	models.DB.Create(&device)

	c.JSON(http.StatusOK, gin.H{"data": device})
}
