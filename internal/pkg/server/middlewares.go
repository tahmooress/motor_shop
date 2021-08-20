package server

import (
	"errors"
	"net/http"

	"github.com/tahmooress/motor-shop/internal/pkg/logger"
)

func recoveryMiddleware(h http.Handler, logger *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t) // nolint: goerr113
				case error:
					err = t
				default:
					err = ErrUnknown
				}
				logger.Error.Printf("panic: %s", err.Error())
				makeResponseError(w, logger, err, http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "PUT")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE")
		w.Header().Add("Access-Control-Allow-Methods", "OPTIONS")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		w.Header().Add("Access-Control-Allow-Headers", "Accept")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Length")
		w.Header().Add("Access-Control-Allow-Headers", "Authorization")
		w.Header().Add("Access-Control-Allow-Headers", "X-CSRF-Token")
		w.Header().Add("Access-Control-Request-Method ", "GET")
		//w.Header().Add("Access-Control-Request-Method ", )
		w.Header().Add("Access-Control-Request-Method ", "OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Authorization")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Accept", "text/plain")
		w.Header().Add("Accept", "*/*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("OptionsPassthrough", "true")
		next.ServeHTTP(w, r)
	})
}
