package web

type (
	CreateProductRequest struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		CategoryRef string `json:"categoryRef"`
		Variants    []struct {
			ID    string  `json:"id"`
			Name  string  `json:"name"`
			Code  string  `json:"code"`
			Price float64 `json:"price"`
		}
	}

	CreateProductResponse struct {
		ID string `json:"id"`
	}
)
