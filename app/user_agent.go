package main

import (
	"strconv"
	"strings"
)

func handle_user_agent(req *Request, res *Response) {
	userAgent := strings.Trim(req.Headers["User-Agent"], "\r\n ")
	res.Status = 200
	res.Message = "OK"
	res.Headers["Content-Type"] = "text/plain"
	res.Headers["Content-Length"] = strconv.Itoa(len(userAgent))
	res.Body = []byte(userAgent)
}
