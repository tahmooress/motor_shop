package dtoupdateequities

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/query"
)

type Request struct {
	ID          models.ID      `json:"id"`
	PayedAmount float64        `json:"payed_amount"`
	NextDueDate query.NullTime `json:"next_due_date"`
}
