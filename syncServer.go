package main

import (
	"fmt"
	"net"
)

// SyncServer is a simple server that will tell all its clients to start playing their song
// at the same time
type SyncServer struct {
	ln        net.Listener
	clients   []net.Conn
	stop      chan byte
	newClient chan byte
}

// NewSyncServer creates a new server listening on "bind"
func NewSyncServer(bind string) (*SyncServer, error) {
	s := &SyncServer{}

	ln, err := net.Listen("tcp", bind)
	if err != nil {
		return nil, err
	}

	s.ln = ln
	s.stop = make(chan byte)
	s.newClient = make(chan byte)

	return s, nil
}

// WaitForClients waits for clients to connect to the server
func (s *SyncServer) WaitForClients() {
	for {
		go s.waitForClient()
		select {
		case <-s.newClient:
			continue
		case <-s.stop:
			return
		}
	}
}

// StopWaiting stops waiting for new clients
func (s *SyncServer) StopWaiting() {
	s.stop <- 1
}

// StartMusic sends the signal to start all the clients
func (s *SyncServer) StartMusic() {
	for _, c := range s.clients {
		go c.Write([]byte{1})
	}
}

func (s *SyncServer) waitForClient() {
	conn, err := s.ln.Accept()
	if err != nil {
		fmt.Println("Error while waiting for a client:")
		panic(err)
	}

	s.newClient <- 1

	go s.handleConnection(conn)
}

func (s *SyncServer) handleConnection(conn net.Conn) {
	fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())
	s.clients = append(s.clients, conn)
}
