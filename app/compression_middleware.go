package main

import (
	"slices"
	"strings"
)

func compression_middleware(req *Request, res *Response, next func()) {
	if req.Headers["Accept-Encoding"] == "" {
		next()
	}

	includes_gzip := strings.Split(req.Headers["Accept-Encoding"], ",")
	for i, v := range includes_gzip {
		includes_gzip[i] = strings.TrimSpace(v)
	}
	if slices.Contains(includes_gzip, "gzip") {
		res.Headers["Content-Encoding"] = "gzip"
	}
	next()
}
