package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoupdateequities"
)

func updateReceivablePartly(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("updateShopReceivable >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		var request dtoupdateequities.Request

		err = json.Unmarshal(r.Req, &request)

		response, err := iUseCases.UpdateReceivablePartly(tokenCTX, &request)
		if err != nil {
			return nil, fmt.Errorf("updateShopReceivable >> %w", err)
		}

		return response, nil
	}
}

