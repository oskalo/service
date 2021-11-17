package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID    `json:"product_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
