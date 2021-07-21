package server

import (
	"context"
	"net/http"
	"regexp"

	"github.com/tahmooress/motor-shop/internal/pkg/logger"
)

type Router struct {
	Path     string
	Method   string
	FN       MiddleFunc
	params   []string
	rootPath string
	pattern  *regexp.Regexp
}

func routers(ctx context.Context, mux *http.ServeMux, routers []*Router, logger *logger.Logger) {
	for _, router := range routers {
		err := router.makeParamsAndPth()
		if err != nil {
			panic(err)
		}
	}

	mux.HandleFunc("/", routeHandler(ctx, routers, logger))
}
