package main

import (
	"fmt"
	"net"

	"github.com/Crystalix007/adversarial-dns-server/dns/message"
	"go.uber.org/zap"
)

type Server struct {
	LAddr net.UDPAddr
	Log   *zap.Logger
}

func (s *Server) Start() {
	if s.Log == nil {
		panic("Expected non-nil logger")
	}

	connection, err := net.ListenUDP("udp", &s.LAddr)
	if err != nil {
		fmt.Printf("cmd/server: failed to start UDP socket: %v", err)
		return
	}
	defer connection.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("cmd/server: failed to accept connection")
			continue
		}

		go s.handleDNSReq(connection, addr, buffer[:n])
	}
}

func (s *Server) handleDNSReq(conn *net.UDPConn, addr net.Addr, b []byte) {
	m, err := message.Decode(b)
	if err != nil {
		fmt.Printf("cmd/server: failed to decode UDP packet: %v", err)
		return
	}

	s.Log.Info("Received query", zap.Object("message", m))
}
