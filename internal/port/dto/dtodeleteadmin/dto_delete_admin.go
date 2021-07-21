package dtodeleteadmin

import "github.com/tahmooress/motor-shop/internal/entities/models"

type Request struct {
	AdminID models.ID
}

type Response struct {
	models.Admin
}
