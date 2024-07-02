package internal

import (
	"sort"
	"time"
)

type Item struct {
	Key       string
	Value     any
	UpdatedAt time.Time
}

type LRUCache struct {
	capacity int
	items    []Item
}

func (c *LRUCache) Get(key string) any {
	for i := range c.items {
		if c.items[i].Key == key {
			c.items[i].UpdatedAt = time.Now()
			return c.items[i].Value
		}
	}
	return -1
}

func (c *LRUCache) Set(key string, value any) {
	if len(c.items) < c.capacity {
		c.items = append(c.items, Item{
			Key:       key,
			Value:     value,
			UpdatedAt: time.Now(),
		})
	} else {
		sort.Slice(c.items, func(i, j int) bool {
			return c.items[j].UpdatedAt.After(c.items[i].UpdatedAt)
		})
		c.items[0] = Item{
			Key:       key,
			Value:     value,
			UpdatedAt: time.Now(),
		}
	}
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
	}
}
