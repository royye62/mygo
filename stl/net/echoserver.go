package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	
	l, err := net.Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		fmt.Println("Error Listen: ", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listen on 6666")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error Accept: ", err)
			os.Exit(1)
		}

		fmt.Printf("Accept conn %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	
	for {
		io.Copy(conn, conn)
	}
}

