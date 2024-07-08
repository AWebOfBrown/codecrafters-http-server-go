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
	res.Headers["Content-Type"] = "text/plain"
	res.Headers["Content-Length"] = strconv.Itoa(len(responseString))
	res.Body = []byte(responseString)
}
