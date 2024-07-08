package main

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestRequest(t *testing.T) {
	request_with_body := "POST /files/number HTTP/1.1\r\nContent-Type: application/octet-stream\r\nContent-Length: 5\r\n\r\n12345"

	server, client := net.Pipe()

	go func() {
		defer client.Close()
		io.WriteString(client, request_with_body)
	}()

	req, err := NewRequest(server)
	if err != nil {
		t.Errorf("Request could not be serialised")
	}
	if req.Method != "POST" {
		t.Errorf("Method of Request is not correct")
	}
	if req.Path != "/files/number" {
		t.Errorf("Path of Request is not correct")
	}

	if string(req.Body) != "12345" {
		t.Errorf("Body is not being parsed")
	}
	fmt.Printf("hi: %s", req.Path)
}
