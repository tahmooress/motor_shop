package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func createShopHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		temp := struct {
			ShopName string `json:"shop_name"`
		}{}

		err = json.Unmarshal(r.Req, &temp)
		if err != nil {
			return nil, fmt.Errorf("createShopHandler >> %w", err)
		}

		response, err := iUseCases.CreateShop(tokenCTX, temp.ShopName)
		if err != nil {
			return nil, fmt.Errorf("createShopHandler >> %w", err)
		}

		return response, nil
	}
}
