package main

import (
	"fmt"
	"os"
	"flag"
	"to-do-list/todolist"
	"github.com/gin-gonic/gin"
)

func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}

func main() {
	dbPtr := flag.Int("db", 1, "Database Type, 1=Memory, 2=MariaDB, 3=MongoDB")
    sizePtr := flag.Int("size", 50, "size of ToDo list")
	flag.Parse()
	r := gin.Default()
	err := todolist.Init(*dbPtr, *sizePtr)
	if err != nil{
		fmt.Printf("Usage: %s \n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	r.GET("/tasks", todolist.GetTasks)
	r.GET("/task/:id", todolist.GetTask)
	r.POST("/task", todolist.AddTask)
	r.DELETE("/task/:id", todolist.DeleteTask)
	r.Run(":8080")
}
