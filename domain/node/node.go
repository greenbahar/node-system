package node

import (
	"github.com/google/uuid"
	"github.com/greenbahar/node-system/utils/logger"
)

const (
	ip = "127.0.0.1"
)

type Node struct {
	Id          uuid.UUID
	Blocks      []Block
	BlockHeight uint64

	NodeIP   string
	NodePort string
}

type Block string

func NewNode() *Node {
	id, err := uuid.NewUUID()
	if err != nil {
		logger.Panic("uuid creation for node failed", err)
	}

	return &Node{
		Id:          id,
		BlockHeight: 0,
		Blocks:      make([]Block, 0),

		NodeIP: ip,
		//NodePort: inMemoryDB.GetDbData().NodePort,
	}
}
