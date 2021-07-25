package http

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func updateShopPayable(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("updateShopReceivable >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		eq, ok := r.Params["equityID"]
		if !ok || eq == nil || eq[0] == "" {
			return nil, fmt.Errorf("updateShopReceivable >>  %w", models.ErrEquityID)
		}

		equityID, err := uuid.Parse(eq[0])
		if err != nil {
			return nil, models.ErrUnknown
		}

		response, err := iUseCases.UpdateShopPayable(tokenCTX, equityID)
		if err != nil {
			return nil, fmt.Errorf("updateShopReceivable >> %w", err)
		}

		return response, nil
	}
}
