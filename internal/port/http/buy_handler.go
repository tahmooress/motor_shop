package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/models"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func buyHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		type Request struct {
			Factor models.Factor `json:"factor"`
			ShopID models.ID     `json:"shop_id"`
		}

		var req Request

		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("buyHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("buyHandler >> %w", err)
		}

		err = json.Unmarshal(r.Req, &req)
		if err != nil {
			return nil, fmt.Errorf("buyHandler() >> json.Unmarshal() >> %w", err)
		}

		respFactor, err := iUseCases.Buy(tokenCTX, req.Factor, req.ShopID)
		if err != nil {
			return nil, fmt.Errorf("buyHandler >> %w", err)
		}

		return respFactor, nil
	}
}
