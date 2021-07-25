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
	login = "/login"

	admin   = "/admin"
	adminID = "/:adminID"

	shops  = "/shops"
	shopID = "/:shopID"

	buy = "/buy"

	sell = "/sell"

	inventory = "/inventory"

	factor       = "/factor"
	factorNumber = "/:factorNumber"

	motor = "/motor"
	pelak = "/:pelak"

	customers  = "/customers"
	customer   = "/customer"
	customerID = "/customerID"

	debts    = "/debts"
	demands  = "/demands"
	equityID = "/:equityID"
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
		{Path: admin + adminID, Method: http.MethodDelete, FN: deleteAdminHandler(ctx, iUseCases)},
		{Path: admin, Method: http.MethodGet, FN: getAdminsHandler(ctx, iUseCases)},
		{Path: shops, Method: http.MethodGet, FN: getShopsHandler(ctx, iUseCases)},
		{Path: shops, Method: http.MethodPost, FN: createShopHandler(ctx, iUseCases)},
		{Path: buy, Method: http.MethodPost, FN: buyHandler(ctx, iUseCases)},
		{Path: sell, Method: http.MethodPost, FN: sellHandler(ctx, iUseCases)},
		{Path: inventory + shopID, Method: http.MethodGet, FN: getShopInventoryHandler(ctx, iUseCases)},
		{Path: motor + pelak, Method: http.MethodGet, FN: getMotorByPelakHandler(ctx, iUseCases)},
		{Path: customers, Method: http.MethodGet, FN: getCustomersHandler(ctx, iUseCases)},
		{Path: customer + customerID, Method: http.MethodGet, FN: getCustomerByIDHandler(ctx, iUseCases)},
		{Path: factor + factorNumber + shopID, Method: http.MethodGet, FN: getFactorByNumberHandler(ctx, iUseCases)},
		{Path: debts + shopID, Method: http.MethodGet, FN: getShopPayables(ctx, iUseCases)},
		{Path: demands + shopID, Method: http.MethodGet, FN: getShopReceivable(ctx, iUseCases)},
		{Path: buy + shopID, Method: http.MethodGet, FN: getBuysHandler(ctx, iUseCases)},
		{Path: sell + shopID, Method: http.MethodGet, FN: getSellsHandler(ctx, iUseCases)},
		{Path: demands + equityID, Method: http.MethodPut, FN: updateShopReceivable(ctx, iUseCases)},
		{Path: debts + equityID, Method: http.MethodPut, FN: updateShopPayable(ctx, iUseCases)},
	}

	serverHTTP, err := server.New(ctx, ip, port, routers, logger)
	if err != nil {
		return nil, fmt.Errorf("new >> server.New >> %w", err)
	}

	return serverHTTP, nil
}
