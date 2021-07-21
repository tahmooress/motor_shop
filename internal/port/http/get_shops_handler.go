package http

import (
	"context"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshops"
)

func getShopsHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtogetshops.Request

		request.Query = r.Query

		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		response, err := iUseCases.GetShopsList(tokenCTX, &request)
		if err != nil {
			return nil, fmt.Errorf("getShopsHandler >> %w", err)
		}

		return response, nil
	}
}
