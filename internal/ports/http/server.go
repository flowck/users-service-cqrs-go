package http

import (
	"context"
	"fmt"
	oapi "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net"
	netHttp "net/http"
	"time"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/common/server"
)

type HTTPServer struct {
	s *netHttp.Server
	r *chi.Mux
}

var _ server.Server = (*HTTPServer)(nil)

func NewServer(ctx context.Context, port int, application *app.App) *HTTPServer {
	router := chi.NewRouter()
	handlers := NewHandlers(application)

	initMiddlewares(router)
	HandlerFromMux(handlers, router)

	return &HTTPServer{
		r: router,
		s: &netHttp.Server{
			Addr:              fmt.Sprintf(":%d", port),
			Handler:           router,
			ReadTimeout:       time.Second * 10,
			ReadHeaderTimeout: time.Second * 10,
			WriteTimeout:      time.Second * 10,
			IdleTimeout:       time.Second * 10,
			BaseContext: func(listener net.Listener) context.Context {
				return ctx
			},
		},
	}
}

func (hs *HTTPServer) Start() {
	go func() {
		log.Println("The http server is starting...")
		if err := hs.s.ListenAndServe(); err != nil {
			log.Printf("The http server has stopped due to: %v", err)
		}
	}()
}

func (hs *HTTPServer) Stop(ctx context.Context) {
	if err := hs.s.Shutdown(ctx); err != nil {
		log.Fatalf("An error occurred trying to shutdown the http server: %v", err)
	}
	log.Println("The http server has been stopped gracefully")
}

func initMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}

	//
	// Global error handler needed due to Swagger validation which ends up sending a text
	// response for unhandled errors
	//
	router.Use(oapi.OapiRequestValidatorWithOptions(swagger, &oapi.Options{
		ErrorHandler: func(w netHttp.ResponseWriter, message string, statusCode int) {
			w.WriteHeader(statusCode)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, message)))
		},
	}))
}
