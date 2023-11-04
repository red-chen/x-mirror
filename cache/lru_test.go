package cache

import (
	"strconv"
	"testing"
)

type CacheValue struct {
	size int
}

func (value *CacheValue) Size() int {
	return value.size
}

func TestNewLRUCache(t *testing.T) {
	cache := NewLRUCache(10)

	for i := 0; i < 15; i++ {
		v := new(CacheValue)
		v.size = 1
		cache.Set(strconv.FormatInt(int64(i), 10), v)
	}

	lenght, size, capacity, _ := cache.Stats()
	if lenght != 10 {
		t.Error("Cache lenght not match", lenght)
	}
	if size != 10 {
		t.Error("Cache size not match", size)
	}
	if capacity != 10 {
		t.Error("Cache capacity not match", capacity)
	}
}

func TestLRUCache_All(t *testing.T) {
	cache := NewLRUCache(1024)
	{
		v := new(CacheValue)
		v.size = 1
		cache.Set("1", v)
	}
	{
		v := new(CacheValue)
		v.size = 1
		cache.Set("2", v)
	}
	lenght, _, _, _ := cache.Stats()
	if lenght != 2 {
		t.Error("Cache lenght not match", lenght)
	}

	cache.Delete("2")

	lenght, _, _, _ = cache.Stats()
	if lenght != 1 {
		t.Error("Cache lenght not match", lenght)
	}

}
