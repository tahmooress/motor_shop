package http

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getCustomerByIDHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("getFactorByNumberHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("getCustomerByIDHandler >> %w", err)
		}

		csID, ok := r.Params["customerID"]
		if !ok || csID == nil || csID[0] == "" {
			return nil, fmt.Errorf("getCustomerByIDHandler >>  %w", models.ErrIDIsNotValid)
		}

		id, err := uuid.Parse(csID[0])
		if err != nil {
			return nil, models.ErrIDIsNotValid
		}

		response, err := iUseCases.GetCustomerByID(tokenCTX, id)
		if err != nil {
			return nil, fmt.Errorf("getCustomerByIDHandler >> %w", err)
		}

		return response, nil
	}
}
