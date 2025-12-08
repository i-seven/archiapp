package controllers

import (
	"backendAf/initializers"
	"backendAf/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create
func CreateProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := initializers.DB.Create(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

// Read All
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := initializers.DB.Preload("ArchiStyles").
		Preload("Ceilings").
		Preload("Garages").
		Preload("RoofDetails").
		Preload("SpecialFeatures").
		Find(&products).Error; err != nil {

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, products)
}

// Read Single
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var p models.Product
	if err := initializers.DB.Preload("ArchiStyles").
		Preload("Ceilings").
		Preload("Garages").
		Preload("RoofDetails").
		Preload("SpecialFeatures").
		First(&p, "pid = ?", id).Error; err != nil {

		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, p)
}

// Update
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var existing models.Product
	if err := initializers.DB.First(&existing, "pid = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input.PID = existing.PID
	initializers.DB.Save(&input)

	c.JSON(200, input)
}

// Delete
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := initializers.DB.Delete(&models.Product{}, "pid = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"deleted": true})
}
