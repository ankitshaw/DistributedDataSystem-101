package main

import (
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	log.Println("Performing Read!")
	buf := make([]byte, 2048)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World\r\n"))
	log.Println("Done Writing, Closing!")
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Waiting for incoming connection!")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connection established!")
		go do(conn)
	}
}
