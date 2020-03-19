package cache

import (
	"github.com/Lukaesebrot/mojango"
	"github.com/zekroTJA/timedmap"
	"strings"
	"time"
)

// Cache represents a player cache
type Cache struct {
	apiClient mojango.Mojango
	lifetime time.Duration
	uuids *timedmap.TimedMap
	names *timedmap.TimedMap
}

// New creates a new player cache
func New(lifetime, interval time.Duration) *Cache  {
	return &Cache{
		apiClient: mojango.New(),
		lifetime: lifetime,
		uuids: timedmap.New(interval),
		names: timedmap.New(interval),
	}
}

// GetUUID returns a cached UUID or fetches a new one and caches it
func (cache *Cache) GetUUID(name string) (string, error) {
	// Return the cached value if it exists
	name = strings.ToLower(name)
	if cache.uuids.Contains(name) {
		return cache.uuids.GetValue(name).(string), nil
	}

	// Fetch a new value using the Mojang API
	uuid, err := cache.apiClient.FetchUUID(name); if err != nil {
		return "", err
	}

	// Put the new value inside the cache and return it
	cache.uuids.Set(name, uuid, cache.lifetime)
	return uuid, nil
}

// GetName returns a cached name or fetches a new one and caches it
func (cache *Cache) GetName(uuid string) (string, error) {
	// Return the cached value if it exists
	if cache.names.Contains(uuid) {
		return cache.names.GetValue(uuid).(string), nil
	}

	// Fetch a new value using the Mojang API
	profile, err := cache.apiClient.FetchProfile(uuid, false); if err != nil {
		return "", err
	}
	name := profile.Name

	// Put the new value inside the cache and return it
	cache.names.Set(uuid, name, cache.lifetime)
	return name, nil
}

// Invalidate clears the entire cache
func (cache *Cache) Invalidate() {
	cache.uuids.Flush()
	cache.names.Flush()
}
