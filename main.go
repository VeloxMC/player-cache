package main

import (
	"flag"
	"github.com/VeloxMC/player-cache/api"
	"github.com/VeloxMC/player-cache/cache"
	"time"
)

func main() {
	// Parse the command line flags
	apiAddress := flag.String("apiAddress", ":3030", "The address of the REST API")
	flag.Parse()

	// Create a new cache object
	cacheInstance := cache.New(15 * time.Minute, 1 * time.Minute)

	// Serve the REST API
	go api.Serve(*apiAddress, cacheInstance)
}
