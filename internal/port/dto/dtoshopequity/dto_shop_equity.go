package dtoshopequity

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
	Equity []models.ShopEquity `json:"data"`
	Meta   query.Meta
}
