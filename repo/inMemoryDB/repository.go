package inMemoryDB

import (
	"github.com/greenbahar/node-system/domain/node"
	"github.com/greenbahar/node-system/utils/dns-server"
	"github.com/greenbahar/node-system/utils/logger"
	"go.uber.org/zap"
)

type InMemoDB interface {
	SetPort(port string)
	GetPort() string
	SetBlock(block node.Block)
	GetBlockHeight() uint64
	GetNodeInfo() *node.Node
}

type inMemoDB struct{}

func New() InMemoDB {
	return &inMemoDB{}
}

/*
	to simulate the node functionality like validating TXs and block creation, we assume that it receives blocks from
	hypothetical minors in a string (hash) format regularly from another goroutine
*/

func (m *inMemoDB) GetNodeInfo() *node.Node {
	newNode := node.NewNode()
	newNode.Id = db.Id
	newNode.BlockHeight = db.BlockHeight
	newNode.Blocks = db.Blocks
	newNode.NodeIP = db.NodeIP
	newNode.NodePort = db.NodePort

	return newNode
}
func (m *inMemoDB) SetPort(port string) {
	db.NodePort = port
	dns_server.SetNodeSeed(&dns_server.DNSSeed{IP: db.NodeIP, Port: db.NodePort})

	logger.Info("gRPC server port submitted into the database and dns-seed",
		zap.String("port", port))
}

func (m *inMemoDB) GetPort() string {
	return db.NodePort
}

// SetBlock calls when node start to download blocks from other nodes
// or when it creates blocks itself in case it is the first node in the network
func (m *inMemoDB) SetBlock(block node.Block) {
	db.mu.Lock()

	db.Blocks = append(db.Blocks, block)
	db.BlockHeight++

	db.mu.Unlock()

	logger.Info("block added to the blockchain",
		zap.String("block", string(block)),
		zap.Uint64("blockNumber", db.BlockHeight))
}

func (m *inMemoDB) GetBlockHeight() uint64 {
	db.mu.Lock()
	blockHeight := db.BlockHeight
	db.mu.Unlock()

	return blockHeight
}

// GetBlocksByIndex returns (h, h+s) blocks
func (m *inMemoDB) GetBlocksByIndex(h, s uint64) []node.Block {
	if db.BlockHeight < h+s {
		s = db.BlockHeight - h // prevent "index out of range" error
	}

	db.mu.Lock()
	results := db.Blocks[h : h+s]
	db.mu.Unlock()

	return results
}
