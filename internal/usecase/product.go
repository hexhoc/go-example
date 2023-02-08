package usecase

import (
	"context"
	"hexhoc/go-examples/internal/entity"
)

type ProductUseCase struct {
	repo ProductRepo
}

func NewProductUseCase(repo ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (p *ProductUseCase) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	return p.repo.FindAll(ctx)
}

func (p *ProductUseCase) GetProductById(ctx context.Context, ID int) (entity.Product, error) {
	return p.repo.FindByID(ctx, ID)
}
