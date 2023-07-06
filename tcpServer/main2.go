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

func main() {

	server := spinServer(":4000")
	server.Start()
}
