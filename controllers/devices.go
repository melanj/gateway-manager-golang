package controllers

import (
	"gatewayManager/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateDeviceInput struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	UID    uint   `json:"uid"`
	Vendor string `json:"vendor"`
	//	DateCreated time.Time `json:"dateCreated"`
	Status  string `json:"status"`
	Gateway uint   `json:"gateway"`
}

// FindDevices GET /api/devices
// Get all Devices
func FindDevices(c *gin.Context) {
	var Devices []models.Device

	gatewayId, byGateway := c.GetQuery("gateway_id")
	if byGateway {
		models.DB.Where("gateway_id = ?", gatewayId).Find(&Devices)
	} else {
		models.DB.Find(&Devices)
	}

	c.JSON(http.StatusOK, Devices)
}

// CreateDevice POST /api/devices
// Create new Device
func CreateDevice(c *gin.Context) {
	// Validate input
	var input CreateDeviceInput
	var gateway models.Gateway
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", fmt.Sprint(input.Gateway)).First(&gateway).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown gateway!"})
		return
	}

	// Create Device
	Device := models.Device{ID: input.ID, UID: input.UID, Vendor: input.Vendor, GatewayId: gateway.ID, Status: input.Status}
	models.DB.Create(&Device)

	c.JSON(http.StatusOK, Device)
}

// FindDevice GET /api/devices/:id
// Find a Device
func FindDevice(c *gin.Context) {
	var Device models.Device

	if err := models.DB.Where("id = ?", c.Param("id")).First(&Device).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, Device)
}

// DeleteDevice DELETE /api/devices/:id
// Delete a Device
func DeleteDevice(c *gin.Context) {
	var Device models.Device
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Device).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&Device)

	c.JSON(http.StatusNoContent, nil)
}
