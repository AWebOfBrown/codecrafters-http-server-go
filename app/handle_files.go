package main

import (
	"os"
	"strconv"
	"strings"
)

func handle_files(req *Request, res *Response) {
	switch req.Method {
	case "GET":
		get_file(req, res)
	case "POST":
		post_file(req, res)
	}
}

func get_file(req *Request, res *Response) {
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
	res.Headers["Content-Type"] = "application/octet-stream"
	res.Headers["Content-Length"] = strconv.Itoa(len(data))
	res.Body = data
}

func post_file(req *Request, res *Response) {
	subStrPath := strings.SplitAfterN(req.Path, "/files/", 2)
	filePath := dir + subStrPath[1]
	err := os.WriteFile(filePath, req.Body, 0644)
	if err != nil {
		panic(err)
	}

	res.Status = 201
	res.Message = "Created"
}
