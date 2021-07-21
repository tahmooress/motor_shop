package usecases

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/port/dto/dtodeleteadmin"
)

func (u *UseCases) DeleteAdmin(ctx context.Context, request *dtodeleteadmin.Request) (*dtodeleteadmin.Response, error) {
	response, err := u.IDatabase.DeleteAdmin(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("deleteAdmin >> %w", err)
	}

	return response, nil
}
