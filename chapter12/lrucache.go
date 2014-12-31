// Problem 12.16

package chapter12

import "container/list"

type Cacheable interface {
	Key() string
}

type LRUCacheItem struct {
	cacheable   Cacheable
	listElement *list.Element
}

type LRUCache struct {
	capacity int
	list     *list.List
	items    map[string]*LRUCacheItem
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		items:    make(map[string]*LRUCacheItem),
	}
}

func (c *LRUCache) Get(key string) Cacheable {
	if item, found := c.items[key]; found {
		c.list.MoveToFront(item.listElement)
		return item.cacheable
	} else {
		return nil
	}
}

func (c *LRUCache) Set(cacheable Cacheable) {
	if item, found := c.items[cacheable.Key()]; found {
		item.cacheable = cacheable
		c.list.MoveToFront(item.listElement)
	} else {
		if c.list.Len() == c.capacity {
			tailItem := c.list.Remove(c.list.Back()).(*LRUCacheItem)
			delete(c.items, tailItem.cacheable.Key())
		}
		item := &LRUCacheItem{cacheable: cacheable}
		item.listElement = c.list.PushFront(item)
		c.items[cacheable.Key()] = item
	}
}

func (c *LRUCache) Erase(key string) bool {
	if _, found := c.items[key]; !found {
		return false
	}
	c.list.Remove(c.items[key].listElement)
	delete(c.items, key)
	return true
}
