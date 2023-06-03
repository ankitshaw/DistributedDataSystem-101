package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	listenAddr string
	listener   net.Listener
	quitCh     chan struct{}
}

func spinServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitCh:     make(chan struct{}),
	}
}

func (server *Server) Start() error {
	listener, err := net.Listen("tcp", server.listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	server.listener = listener

	go server.establishConnection()

	<-server.quitCh

	return nil
}

func (server *Server) establishConnection() {
	for {
		fmt.Println("Accepting New Connections")
		conn, err := server.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println("New Connection:", conn.RemoteAddr())
		go server.read(conn)
	}
}

func (server *Server) read(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		msg := buf[:n]
		fmt.Println(string(msg))
	}
}

// func do(conn net.Conn) {
// 	log.Println("Performing Read!")
// 	buf := make([]byte, 2048)
// 	_, err := conn.Read(buf)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	time.Sleep(5 * time.Second)

// 	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World\r\n"))
// 	log.Println("Done Writing, Closing!")
// 	conn.Close()
// }

func main() {
	// listener, err := net.Listen("tcp", ":2000")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	log.Println("Waiting for incoming connection!")
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println("Connection established!")
	// 	go do(conn)
	// }
	server := spinServer(":4000")
	server.Start()
}
