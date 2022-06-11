package syncEngine

import (
	"github.com/greenbahar/node-system/domain/node"
	grpcNode "github.com/greenbahar/node-system/types/node"
)

const (
	S = 10
	n = 3
)

var processCounter = make(map[uint64]blockByCounter) //blockHeight->numberOfReceivedResponses for that height

type blockByCounter struct {
	counter uint64
	block   *grpcNode.Block
}

type RangeResponseProcessor interface {
	ProcessRange(startHeight uint64, blocks []*grpcNode.Block)
	GetActiveRange() (minHeight uint64, maxHeight uint64)
	StartRangeResponseProcessor()
}

type rangeResponseProcessor struct {
	node node.Service
}

func NewRangeResponseProcessor(node node.Service) RangeResponseProcessor {
	return &rangeResponseProcessor{
		node: node,
	}
}

func (r *rangeResponseProcessor) StartRangeResponseProcessor() {
	var startHeight uint64
	var blockRanges []*grpcNode.Block

	for {
		blocks := <-processChannel

		for h, block := range blocks { //blocks map only contains one element
			startHeight = h
			blockRanges = block
		}
		r.ProcessRange(startHeight, blockRanges)

		for height, blockCounter := range processCounter {
			if height == startHeight {
				// set that block to the blockchain db of the new joined node :)
				block := blockCounter.block
				r.node.SetBlock(node.Block(block.Block))
			}
		}
		w.Done()
	}
}

func (r *rangeResponseProcessor) ProcessRange(startHeight uint64, blocks []*grpcNode.Block) {
	minHeight, maxHeight := r.GetActiveRange()

	for index, val := range blocks {
		if startHeight < minHeight || startHeight > maxHeight {
			continue
		}
		blockCounter, ok := processCounter[startHeight+uint64(index)]
		if !ok {
			processCounter[startHeight+uint64(index)] = blockByCounter{counter: 1, block: val}
		}
		blockCounter.counter++
	}
}

func (r *rangeResponseProcessor) GetActiveRange() (minHeight uint64, maxHeight uint64) {
	blockHeight := r.node.GetBlockHeight()
	return blockHeight + 1, blockHeight + S + 1
}
