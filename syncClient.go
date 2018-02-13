package main

import (
	"net"
	"time"
)

// SyncClient connects to a server and waits for it to start the song
type SyncClient struct {
	conn net.Conn
}

// NewSyncClient creates a new client, ready to connect to the given server
func NewSyncClient(remote string) (*SyncClient, error) {
	c := &SyncClient{}

	conn, err := net.Dial("tcp", remote)
	if err != nil {
		return nil, err
	}

	// Setting the deadline to zero will make the Read() call
	// not timeout
	conn.SetDeadline(time.Time{})

	c.conn = conn

	return c, nil
}

// Wait waits for the server to signal it is time to start the song
func (c *SyncClient) Wait() {
	c.conn.Read([]byte{})
	c.conn.Close()
}
