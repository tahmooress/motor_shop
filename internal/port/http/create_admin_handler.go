package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func createAdminHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var adm models.Admin

		//token, err := getToken(r)
		//if err != nil {
		//	return nil, fmt.Errorf("createAdminHandler >> %w", err)
		//}
		//
		//tokenCTX, err := iUseCases.Authentication(ctx, token)
		//if err != nil {
		//	return nil, fmt.Errorf("createAdminHandler >> %w", err)
		//}

		err := json.Unmarshal(r.Req, &adm)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler() >> json.Unmarshal() >> %w", err)
		}

		if len(adm.Shops) == 0 || adm.Shops == nil {
			return nil, models.ErrAdminAccessibilityEmpty
		}

		respAdmin, err := iUseCases.CreateAdmin(ctx, adm)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		return respAdmin, nil
	}
}
