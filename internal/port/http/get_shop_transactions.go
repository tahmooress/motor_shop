package http

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtotransactions"

	"github.com/google/uuid"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getShopTransactionsHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtotransactions.Request

		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("getTransactionsHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("getTransactionsHandler >> %w", err)
		}

		spID, ok := r.Params["shopID"]
		if !ok || spID == nil || spID[0] == "" {
			return nil, fmt.Errorf("getTransactionsHandler >>  %w", models.ErrIDIsNotValid)
		}

		shopID, err := uuid.Parse(spID[0])
		if err != nil {
			return nil, models.ErrIDIsNotValid
		}

		request.ShopID = shopID
		request.Query = r.Query

		response, err := iUseCases.GetShopTransactions(tokenCTX, &request)
		if err != nil {
			return nil, fmt.Errorf("getTransactionsHandler %w", err)
		}

		return response, nil
	}
}
