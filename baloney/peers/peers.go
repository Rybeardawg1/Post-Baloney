package peers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/anacrolix/dht/v2"
	"github.com/anacrolix/dht/v2/krpc"
)

type MagnetMeta struct {
	InfoHash string
	Name     string
	Trackers []string
	Size     string
}

func FindPeers(parsedData MagnetMeta) dht.QueryResult {
	// Initialize a DHT node
	config := dht.NewDefaultServerConfig()

	// Start DHT server
	server, err := dht.NewServer(config)
	if err != nil {
		log.Fatalf("Failed to start DHT server: %v", err)
	}
	defer server.Close()

	// Give time for bootstrap
	time.Sleep(3 * time.Second)

	// Query the DHT network for peers
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := krpc.IdFromString(parsedData.InfoHash)

	peers := server.GetPeers(ctx, nil, id.Int160(), false, dht.QueryRateLimiting{})

	// Print discovered peers
	fmt.Println("Peers found:")
	for _, peer := range peers.Reply.Q {
		fmt.Println(peer)
	}

	return peers
}
