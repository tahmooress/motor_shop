package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) DeleteAdmin(ctx context.Context, admin models.Admin) (*models.Admin, error) {
	respAdmin, err := u.IDatabase.DeleteAdmin(ctx, admin)
	if err != nil {
		return nil, fmt.Errorf("deleteAdmin >> %w", err)
	}

	return respAdmin, nil
}
