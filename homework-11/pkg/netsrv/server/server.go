package server

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"gocore/homework-11/pkg/crawler"
	"gocore/homework-11/pkg/index"
)

const (
	address  = "0.0.0.0:8000"
	protocol = "tcp4"
)

type Server struct {
	listener   net.Listener
	indexer    index.Storage
	data       []crawler.Document
	handleFunc func(string, []crawler.Document, *index.Storage) []byte
}

func New(
	f func(string, []crawler.Document, *index.Storage) []byte,
	data []crawler.Document,
	indexer *index.Storage,
) (*Server, error) {
	listener, err := net.Listen(protocol, address)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener:   listener,
		indexer:    *indexer,
		data:       data,
		handleFunc: f,
	}, nil
}

func (s *Server) Run() error {
	fmt.Printf("Server listening on %q\n", address)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}

		go s.handler(conn)
	}
}

func (s *Server) Close() error {
	if err := s.listener.Close(); err != nil {
		return err
	}
	return nil
}

func (s *Server) handler(conn net.Conn) {
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(time.Second * 60))

	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		response := s.handleFunc(string(msg), s.data, &s.indexer)

		_, err = conn.Write(response)
		if err != nil {
			return
		}
	}
}
