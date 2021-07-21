package models

import "time"

type Inventory struct {
	ID           ID     `json:"id"`
	FactorNumber string `json:"factor_number"`
	Motor        `json:"motor"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Factor struct {
	ID           ID
	FactorNumber string
	PayedAmount  float64
	TotalAmount  float64
	Motors       []Motor
	Equities     []Equity
	Customer     Customer
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Equity struct {
	ID        ID        `json:"id"`
	FactorID  ID        `json:"factor_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	Customer  Customer  `json:"customer"`
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
	ShopIdentity
	Stocks    []Inventory
	Balance   []Transaction
	Equities  []Equity
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShopIdentity struct {
	ID        ID
	ShopName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
