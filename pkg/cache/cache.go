package cache

import (
	"k8s.io/client-go/tools/cache"
)

// Simple concurrency-safe cache for node utility.
// Plugin needs to pull utility per 5 mins.

type nodeUtilityCache struct {
	// NewTTLStore
	// - thread safe
	// - expire entries with an age > ttl
	// - key string, value interface{}
	store cache.ExpirationCache
}

type UtilityCache interface {
	UpdateAllUtility() error
	GetUtility(nodeName string) int
	UpdateNodeUtility(nodeName string, utility int) error
}

// UpdateAllUtility Update utility from remote API to local cache.
func (u *nodeUtilityCache) UpdateAllUtility() error {
	return nil
}

// GetUtility Get node utility.
func (u *nodeUtilityCache) GetUtility(nodeName string) int {
	return 0
}

// UpdateNodeUtility Update node utility.
func (u *nodeUtilityCache) UpdateNodeUtility(nodeName string, utility int) error {
	return nil
}
