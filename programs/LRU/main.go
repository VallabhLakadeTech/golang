package main

import (
	"fmt"
	"sync"
	"time"
)

type LRUCache struct {
	Lock       sync.Mutex
	CacheNodes map[int]time.Time
	Size       int
}

func NewLRUCache(size int) *LRUCache {
	cacheNodes := make(map[int]time.Time)
	return &LRUCache{CacheNodes: cacheNodes, Size: size}
}

func (lruCache *LRUCache) put(data int) {
	lruCache.Lock.Lock()
	defer lruCache.Lock.Unlock()

	cacheNodes := lruCache.CacheNodes
	if lruCache.Size <= len(cacheNodes) {
		var lru int
		lruTime := time.Now()
		// Check if it already exists. If yes, just update its time
		if _, ok := cacheNodes[data]; ok {
			cacheNodes[data] = lruTime
			return
		}
		// If not then add and find least recently used element to delete
		for key, value := range cacheNodes {
			if lruTime.After(value) {
				lru = key
				lruTime = value
			}
		}
		fmt.Println("Removing LRU data: ", lru)
		delete(cacheNodes, lru)
	}
	cacheNodes[data] = time.Now()
}

func (lruCache *LRUCache) get(data int) {
	lruCache.Lock.Lock()
	defer lruCache.Lock.Unlock()
	if _, ok := lruCache.CacheNodes[data]; !ok {
		fmt.Println("Record not found")
		return
	}
	fmt.Printf("%v last accessed at %v\n", data, lruCache.CacheNodes[data])
	lruCache.CacheNodes[data] = time.Now()
}

func main() {

	lruCache := NewLRUCache(5)
	lruCache.put(10)
	lruCache.put(4)
	fmt.Println(lruCache)
	lruCache.put(3)
	lruCache.put(2)
	fmt.Println(lruCache)
	lruCache.put(1)
	lruCache.put(12)
	lruCache.put(17)
	fmt.Println(lruCache)
	lruCache.get(3)
	fmt.Println(lruCache)
	lruCache.put(15)
	lruCache.get(16)
	fmt.Println(lruCache)
}
