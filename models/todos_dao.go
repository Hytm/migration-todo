package models

import (
	"context"
	"log"

	todos_repo "github.com/hytm/migration-todo/db"
)

func Get() ([]Todo, error) {
	var todos []Todo

	rows, err := todos_repo.Client.Query(context.Background(), "SELECT id, description, status FROM todos;")
	if err != nil {
		log.Println("Error when trying to get Todos")
		return todos, err
	}

	for rows.Next() {
		t := Todo{}
		if err := rows.Scan(&t.ID, &t.Description, &t.Status); err != nil {
			log.Println("Error when parsing Todos")
			return todos, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (t *Todo) GetByID() error {
	result := todos_repo.Client.QueryRow(context.Background(), "SELECT id, description, status FROM todos WHERE id=$1;", t.ID)

	if err := result.Scan(&t.ID, &t.Description, &t.Status); err != nil {
		log.Println("Error when trying to get Todo by ID")
		return err
	}
	return nil
}

func (t *Todo) Save() error {
	var lastInsertID int64
	err := todos_repo.Client.QueryRow(context.Background(), "INSERT INTO todos (description, status) VALUES($1, $2) RETURNING id;", t.Description, t.Status).Scan(&lastInsertID)
	if err != nil {
		log.Println("Error when trying to save todo")
		return err
	}
	t.ID = lastInsertID

	return nil
}

func Clean() error {
	_, err := todos_repo.Client.Exec(context.Background(), "TRUNCATE TABLE todos;")
	return err
}
