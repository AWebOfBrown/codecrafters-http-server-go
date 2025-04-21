package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var dir string

func main() {
	dirPtr := flag.String("directory", "/tmp", "specify directory for files")
	flag.Parse()
	dir = *dirPtr

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go func(conn net.Conn) {
			defaultCloseTimer := time.NewTimer(time.Second * 15)
			for {
				select {
				case <-defaultCloseTimer.C:
					conn.Close()
					return
				default:
					var noRequestError *RequestError
					req, noRequestError := NewRequest(conn)
					if noRequestError != nil {
						if errors.As(err, noRequestError) && noRequestError.Code == NoRequestLine {
							break
						} else {
							r := NewResponse(conn)
							r.Status = 400
							r.Message = "Bad Request"
							r.Send()
							break
						}
					}
					res := NewResponse(conn)
					middleware_handler := NewMiddlewareStack(req, res)
					middleware_handler.Use(response_sender, compression_middleware, router)
					middleware_handler.Run()
					if req.Headers["Connection"] == "close" {
						conn.Close()
					} else {
						defaultCloseTimer.Reset(time.Second * 15)
					}

				}
			}
		}(conn)
	}
}
