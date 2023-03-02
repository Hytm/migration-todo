package main

import (
	"fmt"
	"log"
	"time"

	models "github.com/hytm/migration-todo/models"
)

func main() {
	log.Println("Call Get")
	
	_ = models.Get()
	
	time.Sleep(3 * time.Minute)

	_ = models.Get()
}

func initTodos() []models.Todo {
	var todos []models.Todo
	for i := 0; i < 100; i++ {
		t := models.Todo{
			Description: fmt.Sprintf("Description %d %d", i, time.Now().UnixNano()),
			Status:      "Due",
		}
		todos = append(todos, t)
	}
	return todos
}
