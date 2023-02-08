package repo

import (
	"context"
	"fmt"
	"hexhoc/go-examples/internal/entity"
	"hexhoc/go-examples/pkg/postgres"
)

const _defaultEntityCap = 64

type ProductRepo struct {
	*postgres.Postgres
}

func NewProductRepo(p *postgres.Postgres) *ProductRepo {
	return &ProductRepo{p}
}

func (p *ProductRepo) FindAll(ctx context.Context) ([]entity.Product, error) {
	sql, _, err := p.Builder.
		Select("id", "name", "amount", "price").
		From("products").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ProductRepo - FindAll - r.Builder: %w", err)
	}

	rows, err := p.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("ProductRepo - FindAll - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Product, 0, _defaultEntityCap)
	for rows.Next() {
		e := entity.Product{}
		err = rows.Scan(&e.ID, &e.Name, &e.Amount, &e.Price)
		if err != nil {
			return nil, fmt.Errorf("ProductRepo - FindAll - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (p *ProductRepo) FindByID(ctx context.Context, ID int) (entity.Product, error) {
	//TODO implement me
	panic("implement me")
}
