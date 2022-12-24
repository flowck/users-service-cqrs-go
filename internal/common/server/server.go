package server

import (
	"context"
	netHttp "net/http"
)

type Server interface {
	Start()
	Stop(ctx context.Context)
}

func ContentTypeMiddleware(value string) func(handler netHttp.Handler) netHttp.Handler {
	return func(handler netHttp.Handler) netHttp.Handler {
		return netHttp.HandlerFunc(func(w netHttp.ResponseWriter, r *netHttp.Request) {
			w.Header().Set("Content-Type", value)
			handler.ServeHTTP(w, r)
		})
	}
}
