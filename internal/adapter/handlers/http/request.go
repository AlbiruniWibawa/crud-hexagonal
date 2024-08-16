package http

import "time"

type CreateProductRequest struct {
	ProductName string    `json:"product_name" binding:"required"`
	Stock       uint64    `json:"stock" binding:"required" min:"0"`
	IsDeleted   bool      `json:"is_deleted"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateProductRequest struct {
	ProductName string `json:"product_name"`
	Stock       uint64 `json:"stock"`
}
