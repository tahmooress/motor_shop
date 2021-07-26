package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func createTransactionHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createTransactionHandler >> %w", err)
		}

		var tx models.Transaction

		err = json.Unmarshal(r.Req, &tx)
		if err != nil {
			return nil, fmt.Errorf("createTransactionHandler >> %w", err)
		}

		response, err := iUseCases.CreateTransaction(tokenCTX, tx)
		if err != nil {
			return nil, fmt.Errorf("createTransactionHandler >> %w", err)
		}

		return response, nil
	}
}
