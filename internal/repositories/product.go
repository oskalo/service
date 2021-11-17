package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"github.com/oskalo/service/internal/models"
)

const productTableName = "product"

var ErrProductNotFound = fmt.Errorf("product not found")

type product struct {
	db *sql.DB
}

type ProductRepository interface {
	AddProduct(ctx context.Context, model models.Product) error
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProduct(ctx context.Context, id uuid.UUID) (*models.Product, error)
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &product{db}
}

func (p *product) AddProduct(ctx context.Context, model models.Product) error {

	return nil
}

func (p *product) GetProducts(ctx context.Context) ([]models.Product, error) {
	return nil, nil
}

func (p *product) GetProduct(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	return nil, nil
}
