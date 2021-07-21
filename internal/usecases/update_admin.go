package usecases

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoupdateadmin"
)

func (u *UseCases) UpdateAdmin(ctx context.Context, request *dtoupdateadmin.Request) (*dtoupdateadmin.Response, error) {
	hashedPassword, err := u.generateHashPassword(request.Password)
	if err != nil {
		return nil, fmt.Errorf("updateAdmin >> %w", err)
	}

	admin := &models.Admin{
		ID:       request.ID,
		UserName: request.UserName,
		Password: hashedPassword,
		//Accessibility: request.Accessibility,
	}

	updatedAdmin, err := u.IDatabase.UpdateAdmin(ctx, *admin)
	if err != nil {
		return nil, fmt.Errorf("updateAdmin() >> %w", err)
	}

	return &dtoupdateadmin.Response{
		ID:       updatedAdmin.ID,
		UserName: updatedAdmin.UserName,
		Password: updatedAdmin.Password,
		//Accessibility: updatedAdmin.Accessibility,
	}, nil
}
