package main

func response_serializer(req *Request, res *Response, next func()) {
	res.Send()
	next()
}
