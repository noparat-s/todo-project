package main

import (
	"context"
	"todo/cmd"
	"todo/pkg/server"
	"todo/pkg/server/grpc"
	"todo/pkg/server/rest"
	"todo/pkg/service/todo"

	"github.com/gorilla/mux"
)

func main() {
	ctx := cmd.ContextWithSignal(context.Background())
	service := todo.NewService(ctx, mux.NewRouter())
	r := rest.NewRestServer(ctx, service)
	gs := grpc.NewGRPCServer()
	srv := server.NewServer(r, gs)

	srv.Start(ctx)
	defer srv.Close()
	srv.Wait()
}
