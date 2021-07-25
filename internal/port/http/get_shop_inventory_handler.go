package http

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtogetshopinventory"
)

func getShopInventoryHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtogetshopinventory.Request

		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		spID, ok := r.Params["shopID"]
		if !ok || spID == nil || spID[0] == "" {
			return nil, fmt.Errorf("deleteCardHandler >>  %w", models.ErrShopNotExist)
		}

		shopID, err := uuid.Parse(spID[0])
		if err != nil {
			return nil, models.ErrShopNotExist
		}

		request.ShopID = shopID
		request.Query = r.Query

		response, err := iUseCases.GetShopInventory(tokenCTX, &request)
		if err != nil {
			return nil, fmt.Errorf("getShopInventoryHandler %w", err)
		}

		return response, nil
	}
}
