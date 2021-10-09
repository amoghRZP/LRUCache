package LRUCache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	lruCache := New(5)
	assert.Equal(t, 5, lruCache.capacity)
}

func TestCache_Set(t *testing.T) {
	lruCache := New(3)
	lruCache.Set("Amogh", "Mishra")
	lruCache.Set("Akash", "Agarwal")
	lruCache.Set("Vivek", "Awasthi")
	// Makes Amogh Most recent
	val, err := lruCache.Get("Amogh")
	assert.Nil(t, err)
	assert.Equal(t, "Mishra", val)

	lruCache.Set("Raj", "Pandit")
	_, err = lruCache.Get("Akash")
	assert.NotNil(t, err)
}

func TestCache_Get(t *testing.T) {
	lruCache := New(3)
	lruCache.Set("Amogh", "Mishra")
	lruCache.Set("Akash", "Agarwal")
	lruCache.Set("Vivek", "Awasthi")
	// Makes Amogh Most recent
	val, err := lruCache.Get("Amogh")
	assert.Nil(t, err)
	assert.Equal(t, "Mishra", val)
}
