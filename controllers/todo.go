package controllers

import (
	"fmt"
	"net/http"
	"time"
	"to-do/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Query("activity_group_id")
	todos := []models.Todo{}

	if id == "" {
		if err := db.Raw("select * from todos").Scan(&todos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
			return
		}
	} else {
		if err := db.Raw("select * from todos where activity_group_id = ?", id).Scan(&todos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func GetTodoByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	param := c.Param("id")
	var todos models.Todo

	if err := db.Raw("select * from todos where todo_id = ?", param).Scan(&todos).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	}

	if (todos == models.Todo{}) {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": fmt.Sprintf("Todo with ID %s Not Found", param)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func PostTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todos models.Todo
	var B models.PostTodo_Req

	if err := c.ShouldBindJSON(&B); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null"})
		return
	}

	if B.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null"})
		return
	}
	if B.Activity_Group_Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "activity_group_id cannot be null"})
		return
	}

	if B.Priority == "" {
		B.Priority = "very-high"
	}

	fmt.Println(time.Now)

	if err := db.Raw("insert into todos(title, activity_group_id, is_active, priority, created_at, updated_at) values(?,?,?,?,?,?)", B.Title, B.Activity_Group_Id, B.Is_Active, B.Priority, time.Now(), time.Now()).Scan(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	db.Last(&todos)

	c.JSON(http.StatusCreated, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func UpdateTodoByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var todo models.Todo
	var B models.UpdateTodo_Req

	if err := c.ShouldBindJSON(&B); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null", "err": err.Error()})
		return
	}

	result := db.Exec("update todos set title = ?, priority = ?, is_active = ? where todo_id = ?", B.Title, B.Priority, B.Is_Active, id)

	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found", "status": http.StatusNotFound})
		return
	}

	db.First(&todo, "todo_id = ?", id)

	if (todo == models.Todo{}) {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Todo with ID %s Not Found", id), "status": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todo})
}

func DeleteTodoByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	result := db.Exec("delete from todos where todo_id = ?", id)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Todo with ID %s Not Found", id), "status": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": struct{}{}})
}
