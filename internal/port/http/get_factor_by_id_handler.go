package http

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getFactorByNumberHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		fNumber, ok := r.Params["id"]
		if !ok || fNumber == nil || fNumber[0] == "" {
			return nil, fmt.Errorf("deleteCardHandler >>  %w", models.ErrParams)
		}

		factorNumber, err := uuid.Parse(fNumber[0])
		if err != nil {
			return nil, models.ErrUserNotFound
		}



		return nil, nil
	}
}
