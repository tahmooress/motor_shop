package dtogetshops

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

type Request struct {
	Query server.Query
}

type Response struct {
	Data []models.ShopIdentity `json:"data"`
	Meta query.Meta            `json:"meta"`
}
