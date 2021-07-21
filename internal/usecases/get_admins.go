package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtoadmins"
)

func (u *UseCases) GetAdmins(ctx context.Context, request *dtoadmins.Request) (*dtoadmins.Response, error) {
	response, err := u.IDatabase.GetAdmins(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("getAdmins >> %w", err)
	}

	return response, nil
}
