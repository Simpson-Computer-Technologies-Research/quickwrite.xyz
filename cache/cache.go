package cache

// Import packages
import "sync"

// Cache Data and Mutex
type Cache struct {
	mutex *sync.RWMutex
	data  map[string][]byte
}

// Initalize the Cache
func Init() *Cache {
	return &Cache{
		mutex: &sync.RWMutex{},
		data:  make(map[string][]byte),
	}
}

// The Set() function is used to set a key in the cache
func (cache *Cache) Set(key string, value []byte) {
	// Cache mutex locking / unlocking
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	// Limit the cache size
	if len(cache.data) >= 10000 {
		// Delete the first key in the cache data
		for k := range cache.data {
			delete(cache.data, k)
			break
		}
	}

	// Set the key and value inside the cache
	cache.data[key] = value
}

// The ExistsInData() function is used to check whether
// a key exists in the cache data
func (cache *Cache) ExistsInData(key string) bool {
	// Cache mutex locking / unlocking
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	// Variable "e" is a bool for whether the key exists
	var _, exists = cache.data[key]
	return exists
}

// The Get() function is used to get the synonyms
// from the cache data
func (cache *Cache) Get(key string) []byte {
	// Cache mutex locking / unlocking
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	// Return the synonyms
	return cache.data[key]
}
