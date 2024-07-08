package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Headers map[string]string
	Body    []byte
	Method  string
	Path    string
}

func NewRequest(c net.Conn) (*Request, error) {
	reader := bufio.NewReader(c)

	requestLine, _ := reader.ReadString('\n')
	requestLine = strings.TrimSpace(requestLine)

	splitRequestLine := strings.Split(requestLine, " ")
	if len(splitRequestLine) < 3 {
		return nil, fmt.Errorf("invalid HTTP request")
	}

	method := splitRequestLine[0]
	path := splitRequestLine[1]
	// httpVersion := splitRequestLine[2]

	request := &Request{
		Method:  method,
		Path:    path,
		Headers: make(map[string]string),
	}

	for {
		line, _ := reader.ReadString('\n')
		line = strings.Trim(line, "\n\r")
		if line == "" {
			break
		}

		keyValue := strings.SplitN(line, ":", 2)
		request.Headers[keyValue[0]] = strings.Trim(keyValue[1], " ")
	}

	content_length := request.Headers["Content-Length"]
	if content_length != "" {
		length, err := strconv.Atoi(content_length)
		if err != nil {
			return nil, fmt.Errorf("invalid Content-Length header")
		}

		body := make([]byte, length)
		//todo: handle err
		_, err = io.ReadFull(reader, body)
		request.Body = body
	}

	return request, nil
}
