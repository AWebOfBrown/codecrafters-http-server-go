package main

import (
	"fmt"
	"net"
)

type Response struct {
	Status  int
	Message string
	Body    []byte
	Headers map[string]string
	c       net.Conn
	sent    bool
}

func (r *Response) Serialize() ([]byte, error) {
	msgOrEmpty := ""
	if r.Message != "" {
		msgOrEmpty = fmt.Sprintf(" %s", r.Message)
	}
	lineOne := []byte(fmt.Sprintf("HTTP/1.1 %d%s\r\n", r.Status, msgOrEmpty))

	headers := ""
	if r.Headers != nil {
		for key, value := range r.Headers {
			headers += fmt.Sprintf("%s: %s\r\n", key, value)
		}
	}
	headers += "\r\n"
	lineTwo := []byte(headers)

	//Todo: handle content type
	lineThree := r.Body
	// if err != nil {
	// 	return nil, err
	// }

	fullResponse := append(lineOne, lineTwo...)
	fullResponse = append(fullResponse, lineThree...)
	return fullResponse, nil
}

func (r *Response) Send() {
	serializedResponse, _ := r.Serialize()
	r.c.Write(serializedResponse)
	r.sent = true
}

func NewResponse(c net.Conn) *Response {
	return &Response{
		c:    c,
		sent: false,
	}
}
