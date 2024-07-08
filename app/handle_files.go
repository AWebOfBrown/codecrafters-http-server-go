package main

import (
	"os"
	"strconv"
	"strings"
)

func handle_files(req *Request, res *Response) {
	pathSplit := strings.SplitAfter(req.Path, "/files/")
	filePath := dir + pathSplit[1]

	data, err := os.ReadFile(filePath)
	if err != nil {
		res.Status = 404
		res.Message = "Not Found"
		return
	}

	res.Status = 200
	res.Message = "OK"
	res.Headers = map[string]string{}
	res.Headers["Content-Type"] = "application/octet-stream"
	res.Headers["Content-Length"] = strconv.Itoa(len(data))
	res.Body = data
}
