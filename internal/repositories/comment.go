package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"github.com/oskalo/service/internal/models"
)

const commentTableName = "comment"

type comment struct {
	db *sql.DB
}

type CommentRepository interface {
	AddComment(ctx context.Context, model models.Comment) error
	GetCommentForProduct(ctx context.Context, productID uuid.UUID) ([]models.Comment, error)
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &comment{
		db,
	}
}

func (c *comment)AddComment(ctx context.Context, model models.Comment) error {
	return nil
}

func (c *comment)GetCommentForProduct(ctx context.Context, productID uuid.UUID) ([]models.Comment, error) {
	return nil, nil
}
