package controllers

import (
	"github.com/gin-gonic/gin"
	"openrdv-server/models"
	"net/http"
)

func FindDevices(c *gin.Context) {
	var devices []models.Device
	models.DB.Find(&devices)

	c.JSON(http.StatusOK, gin.H{"data": devices})
}
