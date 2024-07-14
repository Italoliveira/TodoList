package handlers

import (
	"App/database"
	"App/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {

	var Tasks []models.Task

	result := database.DB.Find(&Tasks)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, Tasks)
}

func CreateTask(c *gin.Context) {

	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&newTask)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func GetTask(c *gin.Context) {

	taskID := c.Param("id")

	var Task models.Task

	result := database.DB.First(&Task, taskID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, Task)
}

func UpdateTask(c *gin.Context) {

	taskID := c.Param("id")

	var Task models.Task

	result := database.DB.First(&Task, taskID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&Task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = database.DB.Save(&Task)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, Task)
}

func DestroyTask(c *gin.Context) {

	taskID := c.Param("id")

	var Task models.Task

	result := database.DB.Delete(&Task, taskID)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
