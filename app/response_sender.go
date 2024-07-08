package main

func response_sender(req *Request, res *Response, next func()) {
	next()
	res.Send()
}
