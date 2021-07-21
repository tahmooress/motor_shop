package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtoadmin"
)

func (u *UseCases) CreateAdmin(ctx context.Context, request *dtoadmin.Request) (*dtoadmin.Response, error) {
	password, err := u.generateHashPassword(request.Password)
	if err != nil {
		return nil, fmt.Errorf("createAdmin() >> %w", err)
	}

	admin, err := u.IDatabase.CreateAdmin(ctx, request.UserName, password, request.Accessibility)
	if err != nil {
		return nil, fmt.Errorf("createAdmin() >> %w", err)
	}

	return &dtoadmin.Response{
		Admin: *admin,
	}, nil
}
