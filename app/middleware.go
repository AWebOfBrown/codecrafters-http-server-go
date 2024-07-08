package main

type Middleware func(req *Request, res *Response, next func())

type MiddlewareStack struct {
	middlewares []Middleware
	req         *Request
	res         *Response
}

func NewMiddlewareStack(req *Request, res *Response) *MiddlewareStack {
	var middlewares []Middleware
	return &MiddlewareStack{
		middlewares: middlewares,
		req:         req,
		res:         res,
	}
}

func (ms *MiddlewareStack) Use(m ...Middleware) *MiddlewareStack {
	ms.middlewares = append(ms.middlewares, m...)
	return ms
}

func (ms *MiddlewareStack) Run() {
	middlewareQty := len(ms.middlewares)
	var next func(currentIndex int)

	next = func(currentIndex int) {
		if currentIndex == middlewareQty {
			return
		}
		middlewareToCall := ms.middlewares[currentIndex]
		nextIndex := currentIndex + 1
		middlewareToCall(ms.req, ms.res, func() {
			next(nextIndex)
		})
	}
	next(0)
}
