package port

import (
	"context"

	"crud-hexagonal/internal/core/domain"
)

type ProductsRepository interface {
	CreateProducts(ctx context.Context, product *domain.Products) (*domain.Products, error)
	GetByID(ctx context.Context, id uint64) (*domain.Products, error)
	UpdateProducts(ctx context.Context, product *domain.Products) error
	DeleteProducts(ctx context.Context, id uint64) error
	ListProducts(ctx context.Context) ([]*domain.Products, error)
	CountProducts(ctx context.Context) (int64, error)
}

type ProductsService interface {
	CreateProduct(ctx context.Context, product *domain.Products) (*domain.Products, error)
	GetByID(ctx context.Context, id uint64) (*domain.Products, error)
	UpdateProduct(ctx context.Context, product *domain.Products) error
	DeleteProduct(ctx context.Context, id uint64) error
	ListProducts(ctx context.Context) ([]*domain.Products, error)
}
