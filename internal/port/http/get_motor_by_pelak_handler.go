package http

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getMotorByPelakHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		pelakNumber, ok := r.Params["pelak"]
		if !ok || pelakNumber == nil || pelakNumber[0] == "" {
			return nil, fmt.Errorf("deleteCardHandler >>  %w", models.ErrParams)
		}

		respMotor, err := iUseCases.GetMotorByPelakNumber(tokenCTX, pelakNumber[0])
		if err != nil {
			return nil, fmt.Errorf("deleteAdminHandler >> %w", err)
		}

		return respMotor, nil
	}
}
