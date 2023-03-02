package models

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
	todos_repo "github.com/hytm/migration-todo/db"
)

func Get() []Todo {
	var todos []Todo
	var crdbTodos []CRDBTodo

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		rows, err := todos_repo.Client.Query(context.Background(), "SELECT id, description, status FROM todos;")
		if err != nil {
			log.Println("Error when trying to get Todos")
			return
		}

		for rows.Next() {
			t := Todo{}
			if err := rows.Scan(&t.ID, &t.Description, &t.Status); err != nil {
				log.Println("Error when parsing Todos")
				return
			}
			todos = append(todos, t)
		}
	}()

	go func() {
		defer wg.Done()
		rows, err := todos_repo.CRDB.Query(context.Background(), "SELECT id, description, status FROM todos;")
		if err != nil {
			log.Println("Error when trying to get Todos")
			return
		}

		for rows.Next() {
			t := CRDBTodo{}
			if err := rows.Scan(&t.ID, &t.Description, &t.Status); err != nil {
				log.Println("Error when parsing Todos")
				return
			}
			crdbTodos = append(crdbTodos, t)
		}
	}()
	wg.Wait()

	if len(todos) != len(crdbTodos) {
		go migrate(todos)
	} else {
		log.Println("you can stop accessing Postgres from now")
	}

	return todos
}

func migrate(todos []Todo) {
	log.Println("migrating todos")
	for _, t := range todos {
		imp := CRDBTodo{
			Description: t.Description,
			Status: t.Status,
		}
		if err := imp.Save(); err != nil {
			log.Println("can't save Todo to Cockroach: ", err)
		}
	}
	log.Println("migration done")
}

func (t *Todo) GetByID() error {
	result := todos_repo.Client.QueryRow(context.Background(), "SELECT id, description, status FROM todos WHERE id=$1;", t.ID)

	if err := result.Scan(&t.ID, &t.Description, &t.Status); err != nil {
		log.Println("Error when trying to get Todo by ID")
		return err
	}
	return nil
}

func GetTotal() (int, error) {
	var wg sync.WaitGroup
	var pgCount, crdbCount int
	
	wg.Add(2)
	go func() {
		pgResult := todos_repo.Client.QueryRow(context.Background(), "SELECT COUNT(*) FROM todos;")
		if err := pgResult.Scan(&pgCount); err != nil {
			log.Println("Error when trying to count Todos")
		}
		wg.Done()
	}()
	go func() {
		crdbResult := todos_repo.CRDB.QueryRow(context.Background(), "SELECT COUNT(*) FROM todos;")
		if err := crdbResult.Scan(&crdbCount); err != nil {
			log.Println("Error when trying to count Todos")
		}
		wg.Done()
	}()
	wg.Wait()

	log.Println("Counting on PG: ", pgCount)
	log.Println("Counting on CRDB: ", crdbCount)

	return pgCount, nil
}

func (t *CRDBTodo) Save() error {
	var lastInsertID uuid.UUID
	err := todos_repo.CRDB.QueryRow(context.Background(), "INSERT INTO todos (description, status) VALUES($1, $2) RETURNING id;", t.Description, t.Status).Scan(&lastInsertID)
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
