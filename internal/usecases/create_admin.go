package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) CreateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error) {
	password, err := u.generateHashPassword(admin.Password)
	if err != nil {
		return nil, fmt.Errorf("createAdmin() >> %w", err)
	}

	admin.Password = password

	respAdmin, err := u.IDatabase.CreateAdmin(ctx, admin)
	if err != nil {
		return nil, fmt.Errorf("createAdmin() >> %w", err)
	}

	return respAdmin, nil
}
