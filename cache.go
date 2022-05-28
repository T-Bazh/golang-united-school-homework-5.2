package cache

import "time"

type Item struct {
	value             string
	hasExpirationTime bool
	expirationTime    time.Time
}

type Cache struct {
	data map[string]Item
}

func NewCache() *Cache {
	c := &Cache{
		data: make(map[string]Item),
	}
	return c
}

func (c *Cache) Get(key string) (string, bool) {
	item, ok := c.data[key]
	if !ok {
		return "", ok
	}
	if item.hasExpirationTime {
		timeNow := time.Now()
		expired := timeNow.After(item.expirationTime)
		if expired {
			delete(c.data, key)
			return "", false
		} else {
			return item.value, true
		}
	}
	return item.value, true
}

func (c *Cache) Put(key, value string) {
	item := Item{
		value: value,
	}
	c.data[key] = item
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	item := Item{
		value:             value,
		hasExpirationTime: true,
		expirationTime:    deadline,
	}
	c.data[key] = item
}
