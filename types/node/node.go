package node

import (
	"context"
	"github.com/greenbahar/node-system/repo/inMemoryDB"
)

const S = 10

type Server struct{}

func (s *Server) RangeBlockRequest(ctx context.Context, query *ActiveRangeRequest) (*ActiveRangeResponse, error) {
	startBlock := query.StartHeight
	response := &ActiveRangeResponse{}

	blockchain := inMemoryDB.GetDbData().Blocks //[startBlock:startBlock+ S +1]
	for index, val := range blockchain {
		if uint64(index) < startBlock {
			continue
		}

		if uint64(index) > startBlock+S+1 {
			break
		}

		block := string(val)
		//TODO: change the ActiveRangeResponse.Blocks model
		// from "&Block{BlockHeight: uint64(index), Block: block}" into "&Block{Block: block}"
		// because the returned blocks in the array are consecutive and there's no need to label their heights.
		// e.g 10 consecutive blocks in the array started from height 50
		response.Blocks = append(response.Blocks, &Block{BlockHeight: uint64(index), Block: block})
	}

	return response, nil
}

func (s *Server) mustEmbedUnimplementedQueryServer() {}
