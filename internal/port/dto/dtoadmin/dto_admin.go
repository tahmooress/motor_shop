package dtoadmin

import "github.com/tahmooress/motor-shop/internal/entities/models"

type Request struct {
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	Accessibility []models.ID `json:"accessibility"`
}

type Response struct {
	Admin models.Admin
}
