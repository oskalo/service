package comments

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type createCommentRequest struct {
	ProductID string
	Content   string    `json:"content"`
}

type getCommentsRequest struct {
	ProductID string
}

func parseGetCommentsRequest(vars map[string]string) getCommentsRequest {
	return getCommentsRequest{
		ProductID: vars["productID"],
	}
}

func parseCreateCommentRequest(vars map[string]string, body io.Reader) (createCommentRequest, error) {
	ccr := createCommentRequest{
		ProductID: vars["productID"],
	}

	decoder := json.NewDecoder(body)
	err := decoder.Decode(&ccr)
	if err != nil {
		return ccr, err
	}

	return ccr, nil
}

func validate(ccr createCommentRequest) error {
	_, err := uuid.Parse(ccr.ProductID)
	if err != nil {
		return err
	}

	return nil
}
