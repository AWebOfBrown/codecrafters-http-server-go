package main

import (
	"io"
	"net"
	"testing"
)

func TestWritingFile(t *testing.T) {
	request_with_body := "POST /files/pear_apple_pear_grape HTTP/1.1\r\nContent-Type: application/octet-stream\r\nContent-Length: 64\r\n\r\napple grape raspberry pineapple orange blueberry pear strawberry"

	server, client := net.Pipe()

	go func() {
		defer client.Close()
		io.WriteString(client, request_with_body)
	}()

	go func(conn net.Conn) {
		defer conn.Close()
		req, _ := NewRequest(conn)
		res := NewResponse(conn)
		router(req, res)
		res.Send()
	}(server)
}
