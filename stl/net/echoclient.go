package main

import (
	"fmt"
//	"io"
	"net"
	"os"
	"sync"
)

func main() {
	
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("Error Connecting: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to ...")

	var wg sync.WaitGroup
	wg.Add(2)

	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	_, e := conn.Write([]byte("hello\r\n"))	
	if e != nil {
		fmt.Println("Error to send, err: ", e.Error())
	}
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read, err: ", err)
		return 
	}
	fmt.Println(string(buf[:len-1]))
}

