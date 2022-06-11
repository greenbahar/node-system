package syncEngine

import (
	"fmt"
	"github.com/greenbahar/node-system/domain/node"
	"github.com/greenbahar/node-system/repo/grpcClient"
	grpcNode "github.com/greenbahar/node-system/types/node"
	dnsServer "github.com/greenbahar/node-system/utils/dns-server"
	"github.com/greenbahar/node-system/utils/logger"
	"sync"
	"time"
)

const average_number_of_nodes_in_network = 1000

var (
	w              sync.WaitGroup
	processChannel = make(chan map[uint64][]*grpcNode.Block, average_number_of_nodes_in_network)
	//requestResultsStorage = make([]grpcNode.Block, 0)
)

type RangeRequestService interface {
	StartSyncEngine()
	RangeBlockRequestCalls(blockHeight uint64)
}

type rangeRequestService struct {
	node node.Service
	repo Repository
}

type Repository interface {
	RangeBlockRequest(startHeight uint64, dnsSeed dnsServer.DNSSeed) (*grpcNode.ActiveRangeResponse, error)
}

func NewRangeRequestService(node node.Service, repo grpcClient.ClientService) RangeRequestService {
	return &rangeRequestService{
		node: node,
		repo: repo,
	}
}

func (r *rangeRequestService) StartSyncEngine() {
	//create a goroutine to notice if the rangeBlockProcess is "done" in order to send the next rangeBlockRequest call

	// send rangeBlockRequest to network nodes every 1 seconds to update the blockchain ledger
	ttlTimer := time.NewTicker(10 * time.Second)
	timer := time.NewTicker(1 * time.Second)

	select {
	case <-timer.C:
		w.Add(1)

		blockHeight := r.node.GetNodeInfo().BlockHeight
		r.RangeBlockRequestCalls(blockHeight)

		w.Wait()

	case <-ttlTimer.C:
		// TODO
		// there is no other nodes in the blockchain network. node can start receiving valid blocks from miners
		// set blocks from hypothetical minors to the InMemoDB (from another running goroutine)

	}
}

func (r *rangeRequestService) RangeBlockRequestCalls(startHeight uint64) {
	for _, dnsSeed := range dnsServer.DNSSeeds {
		results, err := r.repo.RangeBlockRequest(startHeight, *dnsSeed)
		if err != nil {
			logger.Error(fmt.Sprintf("no response from node: %s", r.node.GetNodeInfo().Id.String()), err)
		}

		setResultsToStorage(startHeight, results)
	}
}

func setResultsToStorage(startHeight uint64, results *grpcNode.ActiveRangeResponse) {
	result := make(map[uint64][]*grpcNode.Block)
	result[startHeight] = results.Blocks
	processChannel <- result

	//for _, val := range results.Blocks{
	//	requestResultsStorage = append(requestResultsStorage, *val)
	//processChannel <- val
	//}
}
