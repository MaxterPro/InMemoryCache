package main

import (
	"fmt"
)

func main() {

	cache := New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}

type Data struct {
	Value interface{}
}

type cache struct {
	items map[string]Data
}

func New() *cache {
	cache := &cache{
		items: make(map[string]Data),
	}
	return cache
}

func (c *cache) Delete(key string) bool {
	_, exist := c.items[key]
	if !exist {
		return false
	}

	delete(c.items, key)
	return true
}

func (c *cache) Set(key string, value interface{}) {
	c.items[key] = Data{
		Value: value,
	}
}

func (c *cache) Get(key string) interface{} {
	data, exist := c.items[key]
	if !exist {
		return nil
	}

	return data.Value
}
