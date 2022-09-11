package main

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("cmd/server/main: failed to setup logger: %+v", err))
	}

	s := Server{
		LAddr: net.UDPAddr{
			IP:   []byte{},
			Port: 5000,
			Zone: "",
		},
		Log: logger,
	}
	s.Start()
}
