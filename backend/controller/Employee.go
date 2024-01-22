package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/waranyafeen/Projectwaranrat/entity"
	"net/http"
)

// POST /employees
func CreateEmployees(c *gin.Context) {
	var employees entity.Employee

	if err := c.ShouldBindJSON(&employees); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employees})

}

// GET /employees/id
func GetEmployees(c *gin.Context) {
	var employees entity.Employee
	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM employees WHERE id = ?", id).Scan(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employees})

}
