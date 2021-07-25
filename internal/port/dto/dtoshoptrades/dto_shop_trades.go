package dtoshoptrades

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

type Request struct {
	ShopID models.ID `json:"shop_id"`
	server.Query
}

type Response struct {
	Data []models.ShopTrades `json:"data"`
	Meta query.Meta          `json:"meta"`
}
