package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/models"
)

func (u *UseCases) GetMotorByPelakNumber(ctx context.Context, pelakNumber string) (*models.Motor, error) {
	motor, err := u.IDatabase.GetMotorByPelakNumber(ctx, pelakNumber)
	if err != nil {
		return nil, fmt.Errorf("GetMotorByPelakNumber >> %w", err)
	}

	return motor, nil
}
