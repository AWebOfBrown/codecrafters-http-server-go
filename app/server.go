package main

import (
	"flag"
	"fmt"
	"net"
	"os"
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
			defer conn.Close()
			req, _ := NewRequest(conn)
			res := NewResponse(conn)
			middleware_handler := NewMiddlewareStack(req, res)
			middleware_handler.Use(response_sender, compression_middleware, router)
			middleware_handler.Run()
		}(conn)
	}
}
