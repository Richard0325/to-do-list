package main

import (
	"flag"
	"to-do-list/todolist"

	"github.com/gin-gonic/gin"
)

func main() {
	dbPtr := flag.Int("db", 1, "Database Type, 1=Memory, 2=MariaDB, 3=MongoDB")
    sizePtr := flag.Int("size", 50, "size of ToDo list")
	flag.Parse()
	r := gin.Default()
	todolist.Init(*dbPtr, *sizePtr)
	r.GET("/tasks", todolist.GetTasks)
	r.GET("/task/:id", todolist.GetTask)
	r.POST("/task", todolist.AddTask)
	r.DELETE("/task/:id", todolist.DeleteTask)
	r.Run(":8080")
}
