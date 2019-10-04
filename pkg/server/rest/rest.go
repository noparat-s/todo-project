package rest

import (
	"context"
	"log"
	"net/http"
	"todo/pkg/server"
)

// RestServer type
type restServer struct {
	httpServer *http.Server
}

// NewRestServer init rest server
func NewRestServer(ctx context.Context, handler http.Handler) server.EntryPoint {
	svr := &http.Server{
		Addr:    ":3000",
		Handler: handler,
	}
	return &restServer{
		httpServer: svr,
	}
}

func (r *restServer) Start(ctx context.Context) error {
	log.Println("Starting rest")
	return r.httpServer.ListenAndServe()
}

func (r *restServer) ShutDown(ctx context.Context) {
	defer log.Println("stopping rest")
	go func() {
		r.httpServer.Shutdown(ctx)
		<-ctx.Done()
	}()
}
