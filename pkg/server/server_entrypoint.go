package server

import "context"

// EntryPoint type
type EntryPoint interface {
	Start(context.Context) error
	ShutDown(context.Context)
}

// EntryPoints list
type EntryPoints []EntryPoint
