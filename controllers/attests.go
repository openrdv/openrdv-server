package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openrdv-server/models"
)

type CreateAttestInput struct {
	DeviceID uint `json:"device_id"`
}

func FindAttests(c *gin.Context) {
	var attests []models.Attest
	models.DB.Find(&attests)

	c.JSON(http.StatusOK, gin.H{"data": attests})
}

func CreateAttest(c *gin.Context) {
	var input CreateAttestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var device models.Device
	models.DB.First(&device, input.DeviceID)

	// Create attest
	attest := models.Attest{Device: device}
	models.DB.Create(&attest)

	c.JSON(http.StatusOK, gin.H{"data": attest})
}
