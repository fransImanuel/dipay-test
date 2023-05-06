package server

import (
	"to-do/controllers"
	"to-do/db"

	"github.com/gin-gonic/gin"
)

func NewRouter(mysql *db.MysqlConn) *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", mysql.GetDBInstance()) //set mysql db instance
	})

	//Activity
	r.GET("/activity-groups", controllers.GetAllActivity)
	r.GET("/activity-groups/:id", controllers.GetActivityByID)
	r.POST("/activity-groups", controllers.PostActivity)
	r.PATCH("/activity-groups/:id", controllers.UpdateActivityByID)
	r.DELETE("/activity-groups/:id", controllers.DeleteActivityByID)

	//Todo
	r.GET("/todo-items", controllers.GetAllTodo)
	r.GET("/todo-items/:id", controllers.GetTodoByID)
	r.POST("/todo-items", controllers.PostTodo)
	r.PATCH("/todo-items/:id", controllers.UpdateTodoByID)
	r.DELETE("/todo-items/:id", controllers.DeleteTodoByID)

	return r

}
