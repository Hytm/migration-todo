package main

import (
	"fmt"
	"log"

	models "github.com/hytm/migration-todo/models"
)

func main() {
	// models.Clean()
	todos := initTodos()
	for _, t := range todos {
		err := t.Save()
		if err != nil {
			log.Printf("Error saving Todo[%s]: %s\n", t.Description, err.Error())
		}
	}

	todos, err := models.Get()
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range todos {
		fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, t.Status)
	}
}

func initTodos() []models.Todo {
	var todos []models.Todo
	for i := 0; i < 100; i++ {
		t := models.Todo{
			Description: fmt.Sprintf("Description %d", i),
			Status:      "Due",
		}
		todos = append(todos, t)
	}
	return todos
}
