package main

import (
	"github.com/greenbahar/node-system/domain/node"
	"github.com/greenbahar/node-system/repo/grpcClient"

	//"github.com/greenbahar/node-system/domain/node"
	"github.com/greenbahar/node-system/domain/syncEngine"
	"github.com/greenbahar/node-system/repo/inMemoryDB"
)

func startApplication(serverPort string) {
	inMemoryDbRepository := inMemoryDB.New()
	inMemoryDbRepository.SetPort(serverPort)
	grpcClientRepository := grpcClient.NewService()

	nodeService := node.NewService(inMemoryDbRepository)

	newRangeRequestService := syncEngine.NewRangeRequestService(nodeService, grpcClientRepository)
	//goroutine to check the state of the blockchain height and start sending request to other node to download the chain
	go newRangeRequestService.StartSyncEngine()

	newSyncEngine := syncEngine.NewRangeResponseProcessor(nodeService)
	go newSyncEngine.StartRangeResponseProcessor()
}
