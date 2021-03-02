package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"net/http"
	"openrdv-server/models"
)

type CreateAttestInput struct {
	UID    string
	Token  string
	Result datatypes.JSON
}

func FindAttests(c *gin.Context) {
	var attests []models.Attest
	models.DB.Find(&attests)

	c.JSON(http.StatusOK, gin.H{"attests": attests})
}

func CreateAttest(c *gin.Context) {
	var input CreateAttestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var device models.Device
	result := models.DB.Where("uid = ?", input.UID).Where("token = ?", input.Token).First(&device)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	// Create attest
	attest := models.Attest{Device: device, Result: input.Result}
	models.DB.Create(&attest)

	c.JSON(http.StatusOK, gin.H{"data": attest})
}
