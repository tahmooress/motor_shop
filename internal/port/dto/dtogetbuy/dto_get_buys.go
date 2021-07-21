package dtogetbuy

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
)

type Request struct {
	Shops []models.ID
}

type Response struct {
	query.Meta `json:"meta"`
	Data       []Inventory `json:"data"`
}

type Inventory struct {
	ShopID  models.ID
	Factors []models.Factor
}
