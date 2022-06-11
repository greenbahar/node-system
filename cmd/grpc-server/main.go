package main

import (
	"github.com/greenbahar/node-system/types/node"
	"github.com/greenbahar/node-system/utils/logger"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func main() {
	logger.Info("gRPC server begin")
	defer logger.Info("gRPC server end")

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		logger.Panic("failed tcp connection", err)
	}

	serverPort := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
	startApplication(serverPort)

	grpcServer := grpc.NewServer()
	s := node.Server{}
	node.RegisterQueryServer(grpcServer, &s)

	if serveErr := grpcServer.Serve(listener); serveErr != nil {
		logger.Panic("serving gRPC server failed", serveErr)
	}
}
