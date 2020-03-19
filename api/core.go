package api

import (
	"github.com/VeloxMC/player-cache/cache"
	routing "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"strings"
)

// cacheInstance holds the cache instance to use for the API
var cacheInstance *cache.Cache

// apiKeys contains all valid API keys
var apiKeys []string

// Serve serves the REST API
func Serve(address string, cache *cache.Cache, validAPIKeys []string) {
	// Create a new router and define the specific endpoints
	router := routing.New()
	router.GET("/uuid/:name", endUUIDHandler)
	router.GET("/name/:uuid", endNameHandler)
	router.POST("/invalidate", endInvalidateHandler)

	// Initialize the cache instance and API keys
	cacheInstance = cache
	apiKeys = validAPIKeys

	// Serve the REST API
	panic(fasthttp.ListenAndServe(address, router.Handler))
}

// validateAuthorization checks whether or not the requester provided a valid API key
func validateAuthorization(ctx *fasthttp.RequestCtx) bool {
	// Check if the requester provided an authorization header
	rawAuthHeader := ctx.Request.Header.Peek("Authorization"); if rawAuthHeader == nil {
		return false
	}

	// Check if the requester specified a valid authorization method
	splitAuthHeader := strings.Split(string(rawAuthHeader), " "); if len(splitAuthHeader) != 2 {
		return false
	}
	if splitAuthHeader[0] != "Bearer" {
		return false
	}

	// Validate the given API key
	for _, key := range apiKeys {
		if key == splitAuthHeader[1] {
			return true
		}
	}
	return false
}
