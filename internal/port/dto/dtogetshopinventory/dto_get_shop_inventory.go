package dtogetshopinventory

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

type Request struct {
	ShopID models.ID
	server.Query
}

type Response struct {
	query.Meta `json:"meta"`
	Data       []models.Inventory `json:"data"`
}
