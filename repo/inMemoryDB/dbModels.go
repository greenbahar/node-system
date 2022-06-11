package inMemoryDB

import (
	"github.com/google/uuid"
	"github.com/greenbahar/node-system/domain/node"
	"sync"
)

// db is the in memory database implemented in each node
var db = &dbModel{}

type dbModel struct {
	mu sync.Mutex

	Id          uuid.UUID
	BlockHeight uint64
	Blocks      []node.Block

	NodeIP   string
	NodePort string
}

func GetDbData() *dbModel {
	db.mu.Lock()
	returnedDb := db
	db.mu.Unlock()

	return returnedDb
}
