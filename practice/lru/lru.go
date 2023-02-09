package lru

import "container/list"

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

type Cache struct {
	maxSize     int64
	currentSize int64
	ll          *list.List
	cache       map[string]*list.Element
	OnEvicted   func(key string, value Value)
}

func New(maxSize int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxSize:     maxSize,
		ll:          list.New(),
		cache:       make(map[string]*list.Element),
		currentSize: 0,
		OnEvicted:   onEvicted,
	}
}

func (c *Cache) Add(key string, value Value) {

	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.currentSize += int64(value.Len()) - int64(kv.value.Len())
		ele.Value = value
	} else {
		entry := &entry{
			key,
			value,
		}
		ele := c.ll.PushFront(entry)
		c.cache[key] = ele
		c.currentSize += int64(len(key) + value.Len())
	}

	for c.maxSize != 0 && c.maxSize < c.currentSize {
		c.RemoveOldest()
	}
}

func (c *Cache) Get(key string) (v Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.currentSize -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
