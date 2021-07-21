package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tahmooress/motor-shop/internal/pkg/logger"
)

func New(ctx context.Context, ip string, port uint64, router []*Router, logger *logger.Logger) (*http.Server, error) {
	if ip == "" || port == 0 || router == nil || len(router) == 0 {
		return nil, ErrIPAndPort
	}

	logger.Info.Printf("http on IP: %s & Port: %d", ip, port)

	mux := http.NewServeMux()

	routers(ctx, mux, router, logger)

	HTTPServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: recoveryMiddleware(jsonMiddleware(mux), logger),
	}

	err := HTTPServer.ListenAndServe()
	if err != nil {
		return nil, fmt.Errorf("http New http error :%w", HTTPServer.ListenAndServe())
	}

	return HTTPServer, nil
}
