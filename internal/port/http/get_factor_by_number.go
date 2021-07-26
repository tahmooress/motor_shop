package http

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getFactorByNumberHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("getFactorByNumberHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("getFactorByNumberHandler >> %w", err)
		}

		fNumber, ok := r.Params["factorNumber"]
		if !ok || fNumber == nil || fNumber[0] == "" {
			return nil, fmt.Errorf("getFactorByNumberHandler >>  %w", models.ErrIDIsNotValid)
		}

		sp, ok := r.Params["shopID"]
		if !ok || sp == nil || sp[0] == "" {
			return nil, fmt.Errorf("getFactorByNumber >>  %w", models.ErrIDIsNotValid)
		}

		spID, err := uuid.Parse(sp[0])
		if err != nil {
			return nil, models.ErrIDIsNotValid
		}

		response, err := iUseCases.GetFactorByNumber(tokenCTX, fNumber[0], spID)
		if err != nil {
			return nil, fmt.Errorf("getFactorByNumberHandler >> %w", err)
		}

		return response, nil
	}
}
