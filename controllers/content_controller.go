package controllers

import (
	"creator-hub-gin/models"
	"creator-hub-gin/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetContents godoc
// @Summary Get all content
// @Description Get all content from dummy data
// @Tags content
// @Produce json
// @Success 200 {array} models.Content
// @Router /contents [get]
func GetContents(c *gin.Context) {
	contents := services.GetAllContents()
	c.JSON(http.StatusOK, contents)
}

// GetContentDetail godoc
// @Summary Get content by ID
// @Description Get content detail by ID
// @Tags content
// @Produce json
// @Param id path int true "Content ID"
// @Success 200 {object} models.Content
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /contents/{id} [get]
func GetContentDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid ID"})
		return
	}
	content := services.GetContentByID(id)
	if content == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Content not found"})
		return
	}
	c.JSON(http.StatusOK, content)
}

// CreateContent godoc
// @Summary Create new content
// @Description Create a new content entry
// @Tags content
// @Accept json
// @Produce json
// @Param content body models.Content true "Content data"
// @Success 201 {object} models.Content
// @Failure 400 {object} models.ErrorResponse
// @Router /contents [post]
func CreateContent(c *gin.Context) {
	var newContent models.Content
	if err := c.ShouldBindJSON(&newContent); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
		return
	}
	created := services.CreateContent(newContent)
	c.JSON(http.StatusCreated, created)
}

// UpdateContent godoc
// @Summary Update existing content
// @Description Update content by ID
// @Tags content
// @Accept json
// @Produce json
// @Param id path int true "Content ID"
// @Param content body models.Content true "Content data"
// @Success 200 {object} models.Content
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /contents/{id} [put]
func UpdateContent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid ID"})
		return
	}
	var updatedContent models.Content
	if err := c.ShouldBindJSON(&updatedContent); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid input"})
		return
	}
	updated := services.UpdateContent(id, updatedContent)
	if updated == nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Content not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteContent godoc
// @Summary Delete content by ID
// @Description Delete content by ID
// @Tags content
// @Produce json
// @Param id path int true "Content ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /contents/{id} [delete]
func DeleteContent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid ID"})
		return
	}
	ok := services.DeleteContent(id)
	if !ok {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Content not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
