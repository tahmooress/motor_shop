package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtoupdateadmin"
)

func updateAdminHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtoupdateadmin.Request

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
			return nil, fmt.Errorf("updateAdminHandler >> json.Unmarshal() >> %w", err)
		}

		response, err := iUseCases.UpdateAdmin(tokenCTX, &request)
		if err != nil {
			return nil, fmt.Errorf("updateAdminHandler >> %w", err)
		}

		return response, nil
	}
}
