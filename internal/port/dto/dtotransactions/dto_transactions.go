package dtotransactions

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
	Data []models.Transaction `json:"data"`
	Meta query.Meta           `json:"meta"`
}
