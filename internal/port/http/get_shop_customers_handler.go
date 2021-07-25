package http

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtocustomers"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getCustomersHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtocustomers.Request

		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		request.Query = r.Query

		response, err := iUseCases.GetCustomers(tokenCTX, &request)
		if err != nil {
			return nil, fmt.Errorf("getCustomersHandler %w", err)
		}

		return response, nil
	}
}
