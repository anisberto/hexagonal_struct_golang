package products

import (
	"errors"
	"github.com/anisberto/hexagonal_struct_golang/internal/categories"
	_ "github.com/anisberto/hexagonal_struct_golang/internal/categories"
)

type Product struct {
	Id       string              `json:"id"`
	Name     string              `json:"name"`
	Category categories.Category `json:"category"`
	Price    float64             `json:"price"`
	IsActive bool                `json:"IsActive"`
}

func _(p *Product) (error, bool) {
	if p != nil || p.Id == "" || p.Name == "" || p.Price == 0.0 {
		return nil, true
	}
	return errors.New("product invalid"), false
}
