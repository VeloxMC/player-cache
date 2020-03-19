package main

import (
	"encoding/json"
	"flag"
	"github.com/VeloxMC/player-cache/api"
	"github.com/VeloxMC/player-cache/cache"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	// Parse the command line flags
	apiAddress := flag.String("apiAddress", ":3030", "The address of the REST API")
	flag.Parse()

	// Read the API keys out of the corresponding file
	file, err := os.Open("keys.json"); if err != nil {
		panic(err)
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file); if err != nil {
		panic(err)
	}
	var apiKeys []string
	err = json.Unmarshal(byteValue, &apiKeys); if err != nil {
		panic(err)
	}

	// Create a new cache object
	cacheInstance := cache.New(15 * time.Minute, 1 * time.Minute)

	// Serve the REST API
	go api.Serve(*apiAddress, cacheInstance, apiKeys)
}
