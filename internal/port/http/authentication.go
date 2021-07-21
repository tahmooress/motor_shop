package http

import (
	"strings"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getToken(request server.RawRequest) (string, error) {
	auth, ok := request.Header["Authorization"]
	if !ok || auth == nil || len(auth) == 0 {
		return "", models.ErrAuthorization
	}

	authValue := strings.Split(auth[0], "bearer ")
	if len(authValue) < 2 {
		return "", models.ErrAuthorization
	}

	return authValue[1], nil
}
