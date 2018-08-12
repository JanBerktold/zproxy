package main

import (
	"errors"
	"net"
	"fmt"
)

type Server struct {
	listener net.Listener
}

func NewServer(listener net.Listener) (*Server, error) {
	if listener == nil {
		return nil, errors.New("net.Listener can not be nil")
	}

	server := &Server{
		listener: listener,
	}

	return server, nil
}

func (s *Server) Listen() error {
	for {
		conn, err := s.listener.Accept()
		if err == nil {
			go s.handle(conn)
		} else {
			fmt.Printf("Bad accept %q\n", err)
		}
	}

	return nil
}

func (s *Server) handle(conn net.Conn) {

}
