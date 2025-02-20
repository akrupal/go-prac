package main

import "log"

type Cache map[string]any

func NewCache() Cache {
	return Cache{}
}

func (m Cache) Get(key string) (any, bool) {
	val := m[key]
	return val, val != nil
}

func (m Cache) Set(key string, val any) {
	m[key] = val
}

func (m Cache) Delete(key string) {
	delete(m, key)
}

func (m Cache) Contains(key string) bool {
	val := m[key]
	return val != nil
}

func (m Cache) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func RunCacheExample() {
	cache := NewCache()

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)

	keys := cache.Keys()
	for k := range keys {
		log.Printf("Key: %v", k)
	}

	a, _ := cache.Get("a")
	log.Printf("a: %v", a)

	b, _ := cache.Get("b")
	log.Printf("b: %v", b)

	z, _ := cache.Get("z")
	log.Printf("z: %v", z)

	cache.Delete("a")
	cache.Delete("z")

	a, exists := cache.Get("a")
	log.Printf("a: %v, exists: %v", a, exists)

	keys = cache.Keys()
	for k := range keys {
		log.Printf("key: %v", k)
	}
}
