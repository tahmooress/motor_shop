package dtocustomers

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

type Request struct {
	server.Query
}

type Response struct {
	Data []models.Customer `json:"data"`
	Meta query.Meta        `json:"meta"`
}
