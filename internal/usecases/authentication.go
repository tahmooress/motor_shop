package usecases

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) Authentication(ctx context.Context, tokenString string) (context.Context, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRETKEY), nil
	})
	if err != nil {
		return nil, fmt.Errorf("authentication >> %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, models.ErrAuthorization
	}

	return context.WithValue(ctx, "props", claims), nil
}
