package models

import (
	"github.com/tahmooress/motor-shop/internal/pkg/query"
)

type Inventory struct {
	ID           ID        `json:"id,omitempty"`
	FactorNumber string    `json:"factor_number,omitempty"`
	Motor        Motor     `json:"motor,omitempty"`
	CreatedAt    query.NullTime `json:"created_at,omitempty"`
	UpdatedAt    query.NullTime `json:"updated_at,omitempty"`
}

type Factor struct {
	ID           ID        `json:"id,omitempty"`
	FactorNumber string    `json:"factor_number,omitempty"`
	PayedAmount  float64   `json:"payed_amount,omitempty"`
	TotalAmount  float64   `json:"total_amount,omitempty"`
	Motors       []Motor   `json:"motors,omitempty"`
	Equities     []Equity  `json:"equities,omitempty"`
	Customer     Customer  `json:"customer,omitempty"`
	CreatedAt    query.NullTime `json:"created_at,omitempty"`
	UpdatedAt    query.NullTime `json:"updated_at,omitempty"`
}

type ShopEquity struct {
	ID           ID        `json:"id,omitempty"`
	CustomerID   ID        `json:"customer_id,omitempty"`
	FactorNumber string    `json:"factor_number,omitempty"`
	Status       string    `json:"status,omitempty"`
	Amount       float64   `json:"amount,omitempty"`
	ClearDate    query.NullTime `json:"clear_date,omitempty"`
	CreatedAt    query.NullTime `json:"created_at,omitempty"`
	UpdatedAt    query.NullTime `json:"updated_at,omitempty"`
}

type Equity struct {
	ID        ID        `json:"id,omitempty"`
	Amount    float64   `json:"amount,omitempty"`
	Status    string    `json:"status,omitempty"`
	DueDate   query.NullTime `json:"due_date,omitempty"`
	CreatedAt query.NullTime `json:"created_at,omitempty"`
	UpdatedAt query.NullTime `json:"updated_at,omitempty"`
}

type Motor struct {
	ID          ID        `json:"id,omitempty"`
	ModelName   string    `json:"model_name,omitempty"`
	PelakNumber string    `json:"pelak_number,omitempty"`
	BodyNumber  string    `json:"body_number,omitempty"`
	Color       string    `json:"color,omitempty"`
	ModelYear   string    `json:"model_year,omitempty"`
	CreatedAt   query.NullTime `json:"created_at,omitempty"`
	UpdatedAt   query.NullTime `json:"updated_at,omitempty"`
}

type Transaction struct {
	ID           ID             `json:"id,omitempty"`
	ShopID       ID             `json:"shop_id,omitempty"`
	FactorNumber string         `json:"factor_number,omitempty"`
	Description  string         `json:"description,omitempty"`
	Subject      string         `json:"subject,omitempty"`
	Type         string         `json:"type,omitempty"`
	Amount       float64        `json:"amount,omitempty"`
	CreatedAt    query.NullTime `json:"created_at,omitempty"`
	UpdatedAt    query.NullTime `json:"updated_at,omitempty"`
}

type Shop struct {
	ID        ID        `json:"id,omitempty"`
	ShopName  string    `json:"shop_name,omitempty"`
	CreatedAt query.NullTime `json:"created_at,omitempty"`
	UpdatedAt query.NullTime `json:"updated_at,omitempty"`
}

type ShopTrades struct {
	FactorNumber string    `json:"factor_number,omitempty"`
	CustomerID   ID        `json:"customer_id,omitempty"`
	TotalAmount  float64   `json:"total_amount,omitempty"`
	PayedAmount  float64   `json:"payed_amount,omitempty"`
	CreatedAt    query.NullTime `json:"created_at,omitempty"`
}
