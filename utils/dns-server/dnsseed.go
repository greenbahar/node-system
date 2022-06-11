package dns_server

import (
	"encoding/json"
	"github.com/greenbahar/node-system/utils/logger"
	"os"
	"sync"
)

var (
	mux      sync.Mutex
	DNSSeeds = make([]*DNSSeed, 0) //list of nodes info: IP-address and ports
)

type DNSSeed struct {
	IP   string
	Port string
}

func init() {
	//TODO create a shared txt file to store dns-seeds then all node instances can read from that. but this service
	// will become a separate micro-service in production design; NO read-write from file and query and issue seed
	// directly to the dns-seed micro-service.

	GetNodeSeeds()
}

func SetNodeSeed(seed *DNSSeed) {
	DNSSeeds = GetNodeSeeds()
	DNSSeeds = append(DNSSeeds, &DNSSeed{IP: seed.IP, Port: seed.Port})

	seeds, err := json.Marshal(DNSSeeds)
	if err != nil {

	}

	mux.Lock()
	writeErr := os.WriteFile("/dnsSeed.txt", seeds, 0644)
	if writeErr != nil {
		logger.Panic("cannot update data to the dns seed server", writeErr)
	}
	mux.Unlock()
}

func GetNodeSeeds() []*DNSSeed {
	mux.Lock()
	seedFile, err := os.ReadFile("/dnsSeed.txt")
	if err != nil {
		logger.Panic("cannot connect to the dns seed server", err)
	}

	if err := json.Unmarshal(seedFile, &DNSSeeds); err != nil {
		logger.Panic("cannot retrieve dns-seeds from dns seed server", err)
	}
	mux.Unlock()

	return DNSSeeds
}
