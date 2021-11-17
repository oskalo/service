package product

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/oskalo/service/internal/repositories"
	"github.com/oskalo/service/pkg/responses"
	"net/http"
)

type Product struct {
	repo repositories.ProductRepository
}

func NewProduct(repo repositories.ProductRepository) *Product {
	return &Product{
		repo,
	}
}

func (c *Product)AddProduct(w http.ResponseWriter, r *http.Request)  {

}

func (c *Product)GetProduct(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()

	p := parseGetProductRequest(mux.Vars(r))

	comments, err := c.repo.GetProduct(ctx, uuid.MustParse(p.ProductID))
	if err == repositories.ErrProductNotFound {
		fmt.Println(responses.JSONErrorResponse(w, http.StatusNotFound, err))
		return
	}
	if err != nil {
		fmt.Println(responses.JSONErrorResponse(w, http.StatusInternalServerError, err))
		return
	}

	fmt.Println(responses.JSONResponse(w, http.StatusOK, comments))
}
