package http

import (
	"context"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

func getSellsHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		//var request dtogetsell.Request
		//
		//token, err := getToken(r)
		//if err != nil {
		//	return nil, fmt.Errorf("createAdminHandler >> %w", err)
		//}
		//
		//tokenCTX, err := iUseCases.Authentication(ctx, token)
		//if err != nil {
		//	return nil, fmt.Errorf("createAdminHandler >> %w", err)
		//}
		//
		//userID, ok := r.Params["id"]
		//if !ok || userID == nil || userID[0] == "" {
		//	return nil, fmt.Errorf("deleteCardHandler >>  %w", models.ErrParams)
		//}
		//
		//id, err := uuid.Parse(userID[0])
		//if err != nil {
		//	return nil, models.ErrUserNotFound
		//}

		return nil, nil
	}
}
