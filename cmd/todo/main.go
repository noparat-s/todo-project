package main

import (
	"context"
	"todo/cmd"
	"todo/pkg/server"
	"todo/pkg/server/rest"

	"github.com/gorilla/mux"
)

func main() {
	ctx := cmd.ContextWithSignal(context.Background())
	handler := mux.NewRouter()
	r := rest.NewRestServer(ctx, handler)
	srv := server.NewServer(r)

	srv.Start(ctx)
	defer srv.Close()
	srv.Wait()
}
