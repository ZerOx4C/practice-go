package main

import (
	"fmt"
	"net"
	"net/http"
)

type Server struct {
	Url     string
	Port    int
	closeCh chan struct{}
}

func (s *Server) Start(handler http.Handler) error {
	s.closeCh = make(chan struct{})

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Port))
	if err != nil {
		return err
	}

	s.Url = "http://" + listener.Addr().String()

	go http.Serve(listener, handler)
	go func() {
		<-s.closeCh
		listener.Close()
	}()

	return nil
}

func (s *Server) Close() {
	close(s.closeCh)
}

func (s *Server) Wait() <-chan struct{} {
	return s.closeCh
}
