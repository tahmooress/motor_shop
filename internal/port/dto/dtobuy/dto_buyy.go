package dtobuy

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"time"
)

type Request struct {
	FactorNumber string          `json:"factor_number"`
	TotalAmount  float64         `json:"total_amount"`
	PayedAmount  float64         `json:"payed_amount"`
	Date         time.Time       `json:"date"`
	ShopID       models.ID       `json:"shop_id"`
	Motors       []models.Motor  `json:"motors"`
	Customer     models.Customer `json:"customer"`
	Equities     []models.Equity `json:"equities"`
}

type Response struct {
	ID           models.ID       `json:"id"`
	FactorNumber string          `json:"factor_number"`
	PayedAmount  float64         `json:"payed_amount"`
	TotalAmount  float64         `json:"total_amount"`
	Motors       []models.Motor  `json:"motors"`
	Equities     []Equity        `json:"equities"`
	Customer     models.Customer `json:"customer"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type Equity struct {
	ID        models.ID `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
