package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
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
		fmt.Println("Connected \n")
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buff := bufio.NewReader(conn)
	reqLine, err := buff.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading connection: ", err.Error())
		os.Exit(1)
	}

	fmt.Printf("req string: %q \n", reqLine)

	split := strings.Split(reqLine, " ")

	addr := split[1]
	fmt.Printf("addr: %q \n", addr)

	var responseContent []byte
	if addr == "/" {
		responseContent = []byte("HTTP/1.1 200 OK\r\n\r\n")
	} else {
		responseContent = []byte("HTTP/1.1 404 Not Found\r\n\r\n")
	}

	fmt.Printf("responseContent: %q", responseContent)

	n, err := conn.Write(responseContent)
	if err != nil {
		fmt.Printf("err: ", err)
	}
	fmt.Printf("bytes: %d \n", n)
	conn.Close()
}
