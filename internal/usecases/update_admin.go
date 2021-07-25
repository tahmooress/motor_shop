package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) UpdateAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error) {
	hashedPassword, err := u.generateHashPassword(admin.Password)
	if err != nil {
		return nil, fmt.Errorf("updateAdmin >> %w", err)
	}

	admin.Password = hashedPassword

	updatedAdmin, err := u.IDatabase.UpdateAdmin(ctx, admin)
	if err != nil {
		return nil, fmt.Errorf("updateAdmin() >> %w", err)
	}

	return updatedAdmin, nil
}
