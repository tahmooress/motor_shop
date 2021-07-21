package http

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtodeleteadmin"
)

func deleteAdminHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtodeleteadmin.Request

		token, err := getToken(r)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		tokenCTX, err := iUseCases.Authentication(ctx, token)
		if err != nil {
			return nil, fmt.Errorf("createAdminHandler >> %w", err)
		}

		userID, ok := r.Params["id"]
		if !ok || userID == nil || userID[0] == "" {
			return nil, fmt.Errorf("deleteCardHandler >>  %w", models.ErrParams)
		}

		id, err := uuid.Parse(userID[0])
		if err != nil {
			return nil, models.ErrUserNotFound
		}

		request.AdminID = id

		response, err := iUseCases.DeleteAdmin(tokenCTX, &request)
		if err != nil {
			return nil, fmt.Errorf("deleteAdminHandler >> %w", err)
		}

		return response, nil
	}
}
