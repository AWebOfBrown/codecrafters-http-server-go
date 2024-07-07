package main

import (
	"strconv"
	"strings"
)

func handleEcho(addr string) string {
	splitAddr := strings.Split(addr, "/")
	return splitAddr[2]
}

func router(req *Request, res *Response) {
	if req.Path == "/" {
		res.Status = 200
		res.Message = "OK"
		res.Headers = map[string]string{
			"Content-Type": "text/plain",
		}
	} else if strings.HasPrefix(req.Path, "/echo/") {
		responseString := handleEcho(req.Path)

		res.Status = 200
		res.Message = "OK"
		res.Headers = map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(responseString)),
		}
		res.Body = []byte(responseString)
	} else if strings.HasPrefix(req.Path, "/user-agent") {
		userAgent := strings.Trim(req.Headers["User-Agent"], "\r\n ")
		res.Status = 200
		res.Message = "OK"
		res.Headers = map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": strconv.Itoa(len(userAgent)),
		}
		res.Body = []byte(userAgent)

	} else {
		{
			res.Status = 404
			res.Message = "Not Found"
			res.Headers = map[string]string{
				"Content-Type": "text/plain",
			}
		}
	}
}
