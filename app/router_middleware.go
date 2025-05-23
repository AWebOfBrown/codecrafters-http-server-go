package main

import (
	"strings"
)

func router(req *Request, res *Response, next func()) {
	if req.Path == "/" {
		res.Status = 200
		res.Message = "OK"
		res.Headers["Content-Type"] = "text/plain"
	} else if strings.HasPrefix(req.Path, "/echo/") {
		handle_echo(req, res)
	} else if strings.HasPrefix(req.Path, "/user-agent") {
		handle_user_agent(req, res)
	} else if strings.HasPrefix(req.Path, "/files/") {
		handle_files(req, res)
	} else {
		res.Status = 404
		res.Message = "Not Found"
		res.Headers["Content-Type"] = "text/plain"
	}
	next()
}
