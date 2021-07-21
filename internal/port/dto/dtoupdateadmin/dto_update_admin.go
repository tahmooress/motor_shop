package dtoupdateadmin

import (
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

type Request struct {
	ID            models.ID   `json:"id,omitempty"`
	UserName      string      `json:"user_name,omitempty"`
	Password      string      `json:"password,omitempty"`
	Accessibility []models.ID `json:"accessibility,omitempty"`
}

type Response struct {
	ID            models.ID   `json:"id,omitempty"`
	UserName      string      `json:"user_name,omitempty"`
	Password      string      `json:"password,omitempty"`
	Accessibility []models.ID `json:"accessibility,omitempty"`
}
