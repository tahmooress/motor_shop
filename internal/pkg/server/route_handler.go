package server

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/tahmooress/motor-shop/internal/pkg/logger"
)

func routeHandler(ctx context.Context, routes []*Router, logger *logger.Logger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		if path[len(path)-1:] != "/" {
			path += "/"
		}

		for _, route := range routes {
			if route.pattern.MatchString(path) && route.Method == req.Method {
				rawReq := makeRequest(res, req, logger)

				for i, param := range route.params {
					rawReq.Params[param] = []string{strings.Split(strings.TrimPrefix(path, route.rootPath), "/")[i]}
				}

				response, err := callFunction(ctx, route.FN, rawReq)
				if err != nil {
					makeResponseError(res, logger, err, http.StatusBadRequest)

					return
				}

				makeResponse(res, response, logger)

				return
			}
		}

		makeResponseError(res, logger, ErrRouteNotFound, http.StatusNotFound)
	}
}

func (r *Router) makeParamsAndPth() error {
	pattern := r.Path
	r.rootPath = r.Path

	if strings.Contains(r.Path, ":") {
		pattern = r.Path[:strings.Index(r.Path, ":")]
		r.rootPath = r.Path[:strings.Index(r.Path, ":")]
	}

	params := make([]string, 0)

	for _, s := range strings.Split(r.Path, "/") {
		if strings.Contains(s, ":") {
			params = append(params, s[1:])
			pattern += `.*/`
		}
	}

	pathPattern, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("http pkg makeParamsAndPth regexpCompile >> %w", err)
	}

	r.pattern = pathPattern
	r.params = params

	return nil
}
