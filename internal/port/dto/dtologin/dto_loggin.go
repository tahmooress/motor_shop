package dtologin

import "github.com/tahmooress/motor-shop/internal/entities/models"

type Request struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Response struct {
	Token string `json:"token"`
	Admin models.Admin `json:"admin"`
}
