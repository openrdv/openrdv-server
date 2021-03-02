package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"openrdv-server/models"
)

type CreateDeviceInput struct {
	UID     string
	UIDType string
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

	// Create device if such UID is not occupied
	var old models.Device
	models.DB.Where("uid = ?", input.UID).First(&old)
	if old.UID == input.UID {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Generate an access token for client requests
	token, err := generateRandomStringURLSafe(64)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	device := models.Device{UID: input.UID, UIDType: input.UIDType, Token: token}
	models.DB.Create(&device)
	c.JSON(http.StatusOK, gin.H{"device": device, "token": token})
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomStringURLSafe(n int) (string, error) {
	b, err := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
