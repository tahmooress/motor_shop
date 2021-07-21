package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtosell"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func sellHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtosell.Request

		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("buyHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("buyHandler >> %w", err)
		}

		err = json.Unmarshal(r.Req, &request)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler() >> json.Unmarshal() >> %w", err)
		}

		response, err := iUseCases.Sell(tokenCTX, &request)

		return response, nil
	}
}
