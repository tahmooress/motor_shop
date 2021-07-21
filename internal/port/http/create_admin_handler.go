package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoadmin"
)

func createAdminHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtoadmin.Request

		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		err = json.Unmarshal(r.Req, &request)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler() >> json.Unmarshal() >> %w", err)
		}

		if len(request.Accessibility) == 0 || request.Accessibility == nil {
			return nil, models.ErrAdminAccessibilityEmpty
		}

		response, err := iUseCases.CreateAdmin(tokenCTX, &request)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		return response, nil
	}
}
