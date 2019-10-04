package server

import (
	"context"
	"sync"
	"time"
)

// Server type
type Server struct {
	entrypoints EntryPoints
	stopChan    chan bool
}

// NewServer create new server
func NewServer(entries ...EntryPoint) *Server {
	return &Server{
		entrypoints: entries,
		stopChan:    make(chan bool, 1),
	}
}

// Start the server
func (s *Server) Start(ctx context.Context) {
	go func() {
		defer s.Close()
		<-ctx.Done()
		s.Stop()
	}()
	for _, e := range s.entrypoints {
		go e.Start(ctx)
	}
}

// Stop Server
func (s *Server) Stop() {
	var wg sync.WaitGroup
	for _, e := range s.entrypoints {
		wg.Add(1)
		go func(entry EntryPoint) {
			ctx := context.Background()
			defer wg.Done()
			entry.ShutDown(ctx)
		}(e)
	}
	wg.Wait()
	s.stopChan <- true
}

// Wait run server
func (s *Server) Wait() {
	<-s.stopChan
}

// Close close all chanels
func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	go func(ctx context.Context) {
		<-ctx.Done()
		if ctx.Err() == context.Canceled {
			return
		} else if ctx.Err() == context.DeadlineExceeded {
			panic("Timeout while stopping traefik, killing instance ✝")
		}
	}(ctx)
	close(s.stopChan)

	cancel()
}
