package http

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getTransactionByIDHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("getTransactionByIDHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("getTransactionByIDHandler >> %w", err)
		}

		tID, ok := r.Params["txID"]
		if !ok || tID == nil || tID[0] == "" {
			return nil, fmt.Errorf("getTransactionByIDHandler >>  %w", models.ErrIDIsNotValid)
		}

		id, err := uuid.Parse(tID[0])
		if err != nil {
			return nil, models.ErrIDIsNotValid
		}

		tx, err := iUseCases.GetTransactionByID(tokenCTX, id)
		if err != nil {
			return nil, fmt.Errorf("getTransactionByIDHandler >> %w", err)
		}

		return tx, nil
	}
}
