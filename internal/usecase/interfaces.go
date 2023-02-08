// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"hexhoc/go-examples/internal/entity"
)

type (
	Product interface {
		GetProductById(ctx context.Context, ID int) (entity.Product, error)
		GetAllProduct(ctx context.Context) ([]entity.Product, error)
	}

	ProductRepo interface {
		FindAll(ctx context.Context) ([]entity.Product, error)
		FindByID(ctx context.Context, ID int) (entity.Product, error)
	}
)
