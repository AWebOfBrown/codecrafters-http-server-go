package main

import (
	"strconv"
	"strings"
)

func handle_echo(req *Request, res *Response) {
	splitAddr := strings.Split(req.Path, "/")
	responseString := splitAddr[2]

	var contentEncoding string = ""
	if req.Headers["Accept-Encoding"] != "" {
		if req.Headers["Accept-Encoding"] == "gzip" {
			contentEncoding = "gzip"
		}
	}

	res.Status = 200
	res.Message = "OK"
	res.Headers = map[string]string{
		"Content-Type":     "text/plain",
		"Content-Length":   strconv.Itoa(len(responseString)),
		"Content-Encoding": contentEncoding,
	}
	res.Body = []byte(responseString)
}
