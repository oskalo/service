package comments

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/oskalo/service/internal/models"
	"github.com/oskalo/service/internal/repositories"
	"github.com/oskalo/service/pkg/responses"
)

type Comment struct {
	repoComment repositories.CommentRepository
	repoProduct repositories.ProductRepository
}

func NewComment(rComment repositories.CommentRepository, rProduct repositories.ProductRepository) *Comment {
	return &Comment{
		repoComment: rComment,
		repoProduct: rProduct,
	}
}

func (c *Comment) AddComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	p, err := parseCreateCommentRequest(mux.Vars(r), r.Body)
	if err != nil {
		fmt.Println(responses.JSONErrorResponse(w, http.StatusBadRequest, err))
		return
	}

	err = validate(p)
	if err != nil {
		fmt.Println(responses.JSONErrorResponse(w, http.StatusBadRequest, err))
		return
	}

	//define a comment to save
	model := models.Comment{
		ID:        uuid.New(),
		ProductID: uuid.MustParse(p.ProductID),
		Content:   p.Content,
	}

	//check if product exist
	_, err = c.repoProduct.GetProduct(ctx, uuid.MustParse(p.ProductID))
	if errors.Is(err, repositories.ErrProductNotFound) {
		fmt.Println(responses.JSONErrorResponse(w, http.StatusNotFound, err))
		return
	}
	if err != nil {
		fmt.Println(responses.JSONErrorResponse(w, http.StatusInternalServerError, err))
		return
	}

	//save comment
	err = c.repoComment.AddComment(ctx, model)
	if err != nil {
		fmt.Println(responses.JSONErrorResponse(w, http.StatusInternalServerError, err))
		return
	}

	fmt.Println(responses.JSONResponse(w, http.StatusCreated, model))
}

func (c *Comment) GetComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	p := parseGetCommentsRequest(mux.Vars(r))

	comments, err := c.repoComment.GetCommentForProduct(ctx, uuid.MustParse(p.ProductID))
	if err != nil {
		fmt.Println(responses.JSONErrorResponse(w, http.StatusInternalServerError, err))
		return
	}

	fmt.Println(responses.JSONResponse(w, http.StatusOK, comments))
}
