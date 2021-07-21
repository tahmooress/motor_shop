package dtogetsell

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
)

type Request struct {
	ShopID models.Shop
}

type Response struct {
	query.Meta `json:"meta"`
	Data       []models.Factor
}
