package controllers

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateGatewayInput struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Serial      string `json:"serial"`
	Name        string `json:"name"`
	IPv4Address string `json:"ipv4Address"`
}

// FindGateways GET /api/gateways
// Get all gateways
func FindGateways(c *gin.Context) {
	var gateways []models.Gateway
	models.DB.Find(&gateways)

	c.JSON(http.StatusOK, gateways)
}

// CreateGateway POST /api/gateways
// Create new gateway
func CreateGateway(c *gin.Context) {
	// Validate input
	var input CreateGatewayInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create gateway
	gateway := models.Gateway{ID: input.ID, Serial: input.Serial, Name: input.Name, IPv4Address: input.IPv4Address}
	models.DB.Create(&gateway)

	c.JSON(http.StatusOK, gateway)
}

// FindGateway GET /api/gateways/:id
// Find a gateway
func FindGateway(c *gin.Context) {
	var gateway models.Gateway

	if err := models.DB.Where("id = ?", c.Param("id")).First(&gateway).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gateway)
}

// DeleteGateway DELETE /api/gateways/:id
// Delete a gateway
func DeleteGateway(c *gin.Context) {
	var gateway models.Gateway
	if err := models.DB.Where("id = ?", c.Param("id")).First(&gateway).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&gateway)

	c.JSON(http.StatusNoContent, gin.H{"data": true})
}
