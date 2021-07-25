package models

import "time"

type Inventory struct {
	ID           ID        `json:"id"`
	FactorNumber string    `json:"factor_number"`
	Motor        Motor     `json:"motor"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Factor struct {
	ID           ID        `json:"id"`
	FactorNumber string    `json:"factor_number"`
	PayedAmount  float64   `json:"payed_amount"`
	TotalAmount  float64   `json:"total_amount"`
	Motors       []Motor   `json:"motors"`
	Equities     []Equity  `json:"equities"`
	Customer     Customer  `json:"customer"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ShopEquity struct {
	ID           ID        `json:"id"`
	CustomerID   ID        `json:"customer_id"`
	FactorNumber string    `json:"factor_number"`
	Status       string    `json:"status"`
	Amount       float64   `json:"amount"`
	ClearDate    time.Time `json:"clear_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Equity struct {
	ID        ID        `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	DueDate   time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Motor struct {
	ID          ID        `json:"id"`
	ModelName   string    `json:"model_name"`
	PelakNumber string    `json:"pelak_number"`
	BodyNumber  string    `json:"body_number"`
	Color       string    `json:"color"`
	ModelYear   string    `json:"model_year"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
	ID                    ID
	ReferenceFactorNumber ID
	Description           string
	Subject               string
	Type                  string
	Amount                float64
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type Shop struct {
	ID        ID        `json:"id"`
	ShopName  string    `json:"shop_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
