package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
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
	defer conn.Close()
	buff := bufio.NewReader(conn)
	reqLine, err := buff.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading connection: ", err.Error())
		os.Exit(1)
	}

	// fmt.Printf("req string: %q \n", reqLine)

	split := strings.Split(reqLine, " ")

	addr := split[1]

	var response *Response
	if addr == "/" {
		response = &Response{
			Status:  200,
			Message: "OK",
			Body:    nil,
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}
	} else if strings.HasPrefix(addr, "/echo/") {
		responseString := handleEcho(addr)
		response = &Response{
			Status:  200,
			Message: "OK",
			Headers: map[string]string{
				"Content-Type":   "text/plain",
				"Content-Length": strconv.Itoa(len(responseString)),
			},
			Body: []byte(responseString),
		}
	} else {
		response = &Response{
			Status:  404,
			Message: "Not Found",
			Headers: map[string]string{
				"Content-Type": "text/plain",
			},
		}
	}

	serialisedResponse, err := response.Serialize()
	fmt.Printf("response: %s", serialisedResponse)

	if err != nil {
		fmt.Errorf("err: %q", err)
	}
	_, err = conn.Write(serialisedResponse)

	if err != nil {
		fmt.Errorf("err: ", err)
	}
}

func handleEcho(addr string) string {
	splitAddr := strings.Split(addr, "/")
	return splitAddr[2]
}
