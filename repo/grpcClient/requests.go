package grpcClient

import (
	"fmt"
	"github.com/greenbahar/node-system/types/node"
	dns_server "github.com/greenbahar/node-system/utils/dns-server"
	"github.com/greenbahar/node-system/utils/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type ClientService interface {
	RangeBlockRequest(startHeight uint64, dnsSeed dns_server.DNSSeed) (*node.ActiveRangeResponse, error)
}

type clientService struct{}

func NewService() ClientService {
	return &clientService{}
}

func (s *clientService) RangeBlockRequest(startHeight uint64, dnsSeed dns_server.DNSSeed) (*node.ActiveRangeResponse, error) {

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", dnsSeed.IP, dnsSeed.Port), grpc.WithInsecure())
	if err != nil {
		logger.Error("did not connect: %s", err)
		return nil, err
	}
	defer conn.Close()

	c := node.NewQueryClient(conn)

	response, responseErr := c.RangeBlockRequest(context.Background(), &node.ActiveRangeRequest{StartHeight: startHeight})
	if responseErr != nil {
		logger.Error("Error when calling RangeBlockRequest: %s", responseErr)
		return nil, responseErr
	}
	log.Printf("response from server: %s", response)
	logger.Info("receive query from node: ")

	return response, nil
}
