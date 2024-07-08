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

func (r *Response) serialiseHead() ([]byte, error) {
	msgOrEmpty := ""
	if r.Message != "" {
		msgOrEmpty = fmt.Sprintf(" %s", r.Message)
	}
	head := []byte(fmt.Sprintf("HTTP/1.1 %d%s\r\n", r.Status, msgOrEmpty))
	return head, nil
}

func (r *Response) serialiseHeaders() ([]byte, error) {
	headersTemplate := ""
	if r.Headers != nil {
		for key, value := range r.Headers {
			headersTemplate += fmt.Sprintf("%s: %s\r\n", key, value)
		}
	}
	headersTemplate += "\r\n"
	headers := []byte(headersTemplate)
	return headers, nil
}

func (r *Response) Send() {
	if r.sent {
		e := fmt.Errorf("sending response body twice")
		panic(e)
	}
	head, _ := r.serialiseHead()
	headers, _ := r.serialiseHeaders()

	serialisedResponse := append(head, headers...)
	serialisedResponse = append(serialisedResponse, r.Body...)
	bytes, _ := r.c.Write(serialisedResponse)
	fmt.Printf("bytes: %d", bytes)
	r.sent = true
}

func NewResponse(c net.Conn) *Response {
	return &Response{
		c:       c,
		sent:    false,
		Headers: map[string]string{},
	}
}
