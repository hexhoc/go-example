// Package entity defines main entities for business logic (services), database mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

type Product struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float64 `json:"price"`
}
