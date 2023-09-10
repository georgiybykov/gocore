package client

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const (
	address  = "0.0.0.0:8000"
	protocol = "tcp4"
)

type Client struct {
	conn net.Conn
}

func New() (*Client, error) {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Search(lexeme string) (string, error) {
	var response string
	server := bufio.NewReader(c.conn)

	fmt.Fprintf(c.conn, "%s\n", lexeme)

HANDLER:
	for {
		bytes, _, err := server.ReadLine()
		if err != nil {
			return "", err
		}

		response = strings.TrimSpace(string(bytes))
		if len(response) == 0 {
			break HANDLER
		}

		fmt.Println(response)
	}
	return response, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
