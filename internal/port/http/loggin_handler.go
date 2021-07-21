package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/port/dto/dtologin"

	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
)

//LoginHandler is responsible for handling login actions
func loginHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, r server.RawRequest) (interface{}, error) {
		var request dtologin.Request

		err := json.Unmarshal(r.Req, &request)
		if err != nil {
			return nil, fmt.Errorf("loginHandler() >> json.unmarshal() >> %w", err)
		}

		response, err := iUseCases.Login(ctx, &request)
		if err != nil {
			return nil, fmt.Errorf("loginHandler >> %w", err)
		}

		return response, nil
	}
}
