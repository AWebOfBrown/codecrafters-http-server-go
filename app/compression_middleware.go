package main

import (
	"bytes"
	"compress/gzip"
	"slices"
	"strconv"
	"strings"
)

func compression_middleware(req *Request, res *Response, next func()) {
	if req.Headers["Accept-Encoding"] == "" {
		next()
		return
	}

	includes_gzip := strings.Split(req.Headers["Accept-Encoding"], ",")
	for i, v := range includes_gzip {
		includes_gzip[i] = strings.TrimSpace(v)
	}
	gzipCompressionRequested := slices.Contains(includes_gzip, "gzip")

	if gzipCompressionRequested == false {
		next()
		return
	}

	next()
	res.Headers["Content-Encoding"] = "gzip"
	var b bytes.Buffer
	encoder := gzip.NewWriter(&b)
	encoder.Write(res.Body)
	encoder.Close()
	res.Body = b.Bytes()
	res.Headers["Content-Length"] = strconv.Itoa(b.Len())

}
