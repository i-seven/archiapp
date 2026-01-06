package controllers

import (
	"backendAf/config"
	"backendAf/models"
	"backendAf/services"
	"backendAf/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Create
func CreateProduct(c *gin.Context) {
	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&p).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

// Read All
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Preload("ArchiStyles").
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
	if err := config.DB.Preload("ArchiStyles").
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
	if err := config.DB.First(&existing, "pid = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	existing.PName = input.PName
	existing.PDescription = input.PDescription
	existing.PlanDetail = input.PlanDetail
	existing.WhatIsIncluded = input.WhatIsIncluded
	existing.WhatIsNotIncluded = input.WhatIsNotIncluded
	existing.Price = input.Price
	existing.Dem = input.Dem
	config.DB.Save(&existing)

	c.JSON(200, existing)
}

// Delete
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Product{}, "pid = ?", id).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"deleted": true})
}

// POST /products/:productId/images
func UploadProductImage(c *gin.Context) {
	productID := c.Param("id")

	file, err := c.FormFile("image")
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "IMAGE_REQUIRED", "image form field required")
		return
	}

	res, err := services.ImageService.UploadProductImage(productID, file)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "UPLOAD_FAILED", err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusCreated, res)
}

// PUT /products/:productId/images/:imageId
func UpdateProductImage(c *gin.Context) {
	productID := c.Param("id")
	imageID := c.Param("imageId")

	file, err := c.FormFile("image")
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, "IMAGE_REQUIRED", "image form field required")
		return
	}

	res, err := services.ImageService.UpdateProductImage(productID, imageID, file)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "UPDATE_FAILED", err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, res)
}

// DELETE /products/:productId/images/:imageId
func DeleteProductImage(c *gin.Context) {
	productID := c.Param("id")
	imageID := c.Param("imageId")

	if err := services.ImageService.DeleteProductImage(productID, imageID); err != nil {
		if err == os.ErrNotExist {
			utils.JSONError(c, http.StatusNotFound, "NOT_FOUND", "image not found")
			return
		}
		utils.JSONError(c, http.StatusInternalServerError, "DELETE_FAILED", err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, gin.H{"deleted": true})
}

// GET /products/:productId/images
func GetProductImages(c *gin.Context) {
	productID := c.Param("id")

	urls, err := services.ImageService.ListProductImages(productID)
	if err != nil {
		utils.JSONError(c, http.StatusInternalServerError, "LIST_FAILED", err.Error())
		return
	}

	utils.JSONSuccess(c, http.StatusOK, gin.H{"images": urls})
}
