package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID
	Name        string `json:"name"`
	Description string `json:"description"`
}
