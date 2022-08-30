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

// The DeleteFirstDataKey() function is used to
// delete the first key of the cache data
func (cache *Cache) DeleteFirstDataKey() {
	// MAKE SURE TO USE A FUNCTION THAT
	// ALREADY HAS THE cache.mutex LOCKING
	// AND UNLOCKING BEFORE CALLING THIS ONE

	// Iterate over the cache data
	for k := range cache.data {

		// Delete the first key
		delete(cache.data, k)

		// Break the loop and return the function
		return
	}
}

// The Set() function is used to set a key in the cache
func (cache *Cache) Set(key string, value []byte) {

	// Cache mutex locking / unlocking
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	// Make sure the cache data size isn't greater than 2MB
	// (80,000 keys)
	//
	// 2,000,000 bytes is 2MB, divide that by 25
	// (the size of each key+value in the cache)
	//
	if len(cache.data) >= 80000 {
		// Delete the first key in the cache data
		cache.DeleteFirstDataKey()
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
	var _, e = cache.data[key]
	return e
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
