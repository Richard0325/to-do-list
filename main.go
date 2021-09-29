package main

import (
	"to-do-list/todolist"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	todolist.Init()
	r.GET("/tasks", todolist.GetTasks)
	r.GET("/task/:id", todolist.GetTask)
	r.POST("/task", todolist.AddTask)
	r.DELETE("/task/:id", todolist.DeleteTask)
	r.Run(":8080")
}
