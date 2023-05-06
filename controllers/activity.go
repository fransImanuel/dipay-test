package controllers

import (
	"fmt"
	"net/http"
	"to-do/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllActivity(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	activities := []models.Activities{}

	if err := db.Raw("select * from activities").Scan(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activities})
}

func GetActivityByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	activity := models.Activities{}

	if err := db.Raw("select * from activities where activity_id = ?", id).Scan(&activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	if (activity == models.Activities{}) {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Activity with ID %s Not Found", id), "status": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activity})
}

func PostActivity(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var activity models.Activities
	var B models.PostActivity_Req

	if err := c.ShouldBindJSON(&B); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null"})
		return
	}

	if err := db.Exec("Insert into activities(title,email) values(?,?)", B.Title, B.Email).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	db.Last(&activity)

	c.JSON(http.StatusCreated, gin.H{"status": "Success", "message": "Success", "data": activity})

}

func UpdateActivityByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var activity models.Activities
	var B models.PostActivity_Req

	if err := c.ShouldBindJSON(&B); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null"})
		return
	}

	if err := db.Exec("update activities set title = ? where activity_id = ?", B.Title, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	db.First(&activity, "activity_id = ?", id)

	if (activity == models.Activities{}) {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Activity with ID %s Not Found", id), "status": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activity})
}

func DeleteActivityByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	result := db.Exec("delete from activities where activity_id = ?", id)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Activity with ID %s Not Found", id), "status": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": struct{}{}})
}
