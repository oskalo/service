package product

type getProductRequest struct {
	ProductID string
}

func parseGetProductRequest(vars map[string]string) getProductRequest {
	return getProductRequest{
		ProductID: vars["productID"],
	}
}
