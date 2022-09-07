package main

import "net"

func main() {
	s := Server{
		LAddr: net.UDPAddr{
			IP:   []byte{},
			Port: 5000,
			Zone: "",
		},
	}
	s.Start()
}
