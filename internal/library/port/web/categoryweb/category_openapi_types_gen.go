// Package categoryweb provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package categoryweb

// Category defines model for Category.
type Category struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// CategoryIdParameter defines model for CategoryIdParameter.
type CategoryIdParameter string

// GetCategoryByIdResponse defines model for GetCategoryByIdResponse.
type GetCategoryByIdResponse Category

// CreateCategoryJSONBody defines parameters for CreateCategory.
type CreateCategoryJSONBody Category

// CreateCategoryJSONRequestBody defines body for CreateCategory for application/json ContentType.
type CreateCategoryJSONRequestBody CreateCategoryJSONBody
