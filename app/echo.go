package main

import (
	"strconv"
	"strings"
)

func handle_echo(req *Request, res *Response) {
	splitAddr := strings.Split(req.Path, "/")
	responseString := splitAddr[2]

	res.Status = 200
	res.Message = "OK"
	res.Headers = map[string]string{
		"Content-Type":   "text/plain",
		"Content-Length": strconv.Itoa(len(responseString)),
	}
	res.Body = []byte(responseString)
}
