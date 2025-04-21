package main

func response_sender(req *Request, res *Response, next func()) {
	next()
	if req.Headers["Connection"] == "close" {
		res.Headers["Connection"] = "close"
	}
	res.Send()
}
