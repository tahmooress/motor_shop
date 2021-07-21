package http

import (
	"context"
	"fmt"
	"github.com/tahmooress/motor-shop/internal/entities/interfaces"
	"github.com/tahmooress/motor-shop/internal/entities/models"
	"github.com/tahmooress/motor-shop/internal/pkg/logger"
	"github.com/tahmooress/motor-shop/internal/pkg/server"
	"net/http"
	"os"
	"strconv"
)

const (
	login      = "/login"
	admin      = "/admin"
	id         = "/:id"
	shops      = "/shops"
	buy        = "/buy"
	sell       = "/sell"
	inventory  = "/inventory"
	factor     = "/factor"
	factorType = "/:type"
)

func New(ctx context.Context, iUseCases interfaces.IUseCases, logger *logger.Logger) (*http.Server, error) {
	ip := os.Getenv("HTTP_IP")
	portEnv := os.Getenv("HTTP_PORT")

	port, err := strconv.ParseUint(portEnv, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("New >> strconv.ParseUint >> %w", err)
	}

	if ip == "" || port == 0 {
		return nil, models.ErrEmptyIPANDPORT
	}

	routers := []*server.Router{

		{Path: login, Method: http.MethodPost, FN: loginHandler(ctx, iUseCases)},
		{Path: admin, Method: http.MethodPost, FN: createAdminHandler(ctx, iUseCases)},
		{Path: admin, Method: http.MethodPut, FN: updateAdminHandler(ctx, iUseCases)},
		{Path: admin + id, Method: http.MethodDelete, FN: deleteAdminHandler(ctx, iUseCases)},
		{Path: admin, Method: http.MethodGet, FN: getAdminsHandler(ctx, iUseCases)},
		{Path: shops, Method: http.MethodGet, FN: getShopsHandler(ctx, iUseCases)},
		{Path: buy, Method: http.MethodPost, FN: buyHandler(ctx, iUseCases)},
		{Path: sell, Method: http.MethodPost, FN: sellHandler(ctx, iUseCases)},
		{Path: inventory + id, Method: http.MethodGet, FN: getShopInventoryHandler(ctx, iUseCases)},

		{Path: buy + id, Method: http.MethodGet, FN: getBuysHandler(ctx, iUseCases)},

		{Path: sell + id, Method: http.MethodGet, FN: getSellsHandler(ctx, iUseCases)},

		{Path: factor + factorType + id, Method: http.MethodGet, FN: getFactorByNumberHandler(ctx, iUseCases)},
	}

	serverHTTP, err := server.New(ctx, ip, port, routers, logger)
	if err != nil {
		return nil, fmt.Errorf("new >> server.New >> %w", err)
	}

	return serverHTTP, nil
}
