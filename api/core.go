package api

import (
	"github.com/VeloxMC/player-cache/cache"
	routing "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// cacheInstance holds the cache instance to use for the API
var cacheInstance *cache.Cache

// Serve serves the REST API
func Serve(address string, cache *cache.Cache) {
	// Create a new router and define the specific endpoints
	router := routing.New()
	router.GET("/uuid/:name", nil)
	router.GET("/name/:uuid", nil)
	router.POST("/invalidate", nil)

	// Initialize the cache instance
	cacheInstance = cache

	// Serve the REST API
	panic(fasthttp.ListenAndServe(address, router.Handler))
}
