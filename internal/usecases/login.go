package usecases

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtologin"

	"github.com/dgrijalva/jwt-go"
)

func (u *UseCases) Login(ctx context.Context, request *dtologin.Request) (*dtologin.Response, error) {
	adminID, err := u.IDatabase.GetAdminIDByUserName(ctx, request.UserName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}

		return nil, fmt.Errorf("useCases >> login() >> %w", err)
	}

	admin, err := u.IDatabase.GetAdminByID(ctx, *adminID)
	if err != nil {
		return nil, fmt.Errorf("useCases >> login() >> %w", err)
	}

	token, err := u.authorization(admin.UserName, admin.Password, request.Password)
	if err != nil {
		return nil, fmt.Errorf("useCases >> login() >> %w", err)
	}

	return &dtologin.Response{
		Token: token,
		Admin: *admin,
	}, nil
}

func (u *UseCases) authorization(userName, dbPassword, password string) (string, error) {
	hashedPass, err := u.generateHashPassword(password)
	if err != nil {
		return "", fmt.Errorf("authorization >> %w", err)
	}

	if hashedPass != dbPassword {
		return "", models.ErrPasswordIsWrong
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userName,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", fmt.Errorf("useCases >> authorization() >> token.SignedString() >> %w", err)
	}

	return tokenString, nil
}
