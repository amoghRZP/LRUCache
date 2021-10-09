package LRUCache

import (
	"errors"
	"time"
)

type Cache struct {
	capacity int
	items    map[string]*CacheItem
}

type CacheItem struct {
	value    string
	lastUsed int64
}

func New(c int) *Cache {
	return &Cache{
		capacity: c,
		items:    make(map[string]*CacheItem),
	}
}

func (c *Cache) Set(key string, val string) {
	// Add value, If key does not exist already
	if _, ok := c.items[key]; !ok {
		lastUsed := int64(time.Now().Nanosecond())
		if len(c.items) == c.capacity {
			var timestamp int64
			var keyToDelete string
			for key, item := range c.items {
				if timestamp == 0 {
					timestamp = item.lastUsed
					keyToDelete = key
				} else if timestamp > item.lastUsed {
					timestamp = item.lastUsed
					keyToDelete = key
				}
			}
			delete(c.items, keyToDelete)
		}
		c.items[key] = &CacheItem{
			value:    val,
			lastUsed: lastUsed,
		}
	}
}

func (c *Cache) Get(key string) (string, error) {
	//Search the key in map
	if item, ok := c.items[key]; ok {
		item.lastUsed = time.Now().UnixNano()
		return item.value, nil
	}
	return "", errors.New("key was not found")
}
