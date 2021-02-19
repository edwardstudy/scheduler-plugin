package cache

import (
	"k8s.io/client-go/tools/cache"
)

// Simple concurrency-safe cache for node utility.
// Plugin needs to pull utility periodically.

type nodeUtilityCache struct {
	// NewThreadSafeStore
	// - thread safe
	// - key string, value interface{}
	store cache.ThreadSafeStore
}

type UtilityCache interface {
	UpdateAllUtility() error
	GetUtility(nodeName string) int
	UpdateNodeUtility(nodeName string, utility int) error
	DeleteNodeUtility(nodeName string) error
}

func NewNodeUtilityCache() UtilityCache {
	return &nodeUtilityCache{}
}

// UpdateAllUtility Update utility from remote API to local cache.
// It should be invoked periodically(per 5 mins).
func (u *nodeUtilityCache) UpdateAllUtility() error {
	// 1. get metrics

	// 2. update local
	// 2.1 alert if errors
	// otherwise, u = a*u(10min) + b*u(1h) +c*u(1d)

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

// UpdateNodeUtility Update node utility.
func (u *nodeUtilityCache) DeleteNodeUtility(nodeName string) error {
	return nil
}
