package todos

import (
	"log"

	"github.com/hytm/migration-todo/db/todos_repo"
)

func (t *Todo) Get() error {
	stmt, err := todos_repo.Client.Prepare("SELECT id, description, status FROM todos WHERE id=$1;")
	if err != nil {
		log.Println("Error when trying to prepare statement: ", err.Error)
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(t.ID)

	if err := result.Scan(&t.ID, &t.Description, &t.Status); err != nil {
		log.Println("Error when trying to get Todo by ID")
		return err
	}
	return nil
}

func (t *Todo) Save() error {
	stmt, err := todos_repo.Client.Prepare("INSERT INTO todos(description, priority, status) VALUES($1, $2, $3) RETURNING id;")
	if err != nil {
		log.Println("Error when trying to prepare statement: ", err)
		return err
	}
	defer stmt.Close()

	var lastInsertID int64
	insertErr := stmt.QueryRow(t.Description, t.Status).Scan(&lastInsertID)
	if insertErr != nil {
		log.Println("Error when trying to save todo")
		return err
	}
	t.ID = lastInsertID

	return nil
}
