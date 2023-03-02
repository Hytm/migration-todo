package models

import "github.com/google/uuid"

type Todo struct {
	ID          int  `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CRDBTodo struct {
	ID          uuid.UUID  `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
