package main

import (
	"fmt"
)

type Store interface {
	Store(key, value string) error
}

type RedisStore struct{}

func (r RedisStore) Store(key, value string) error {
	fmt.Printf("Storing (%s,%s) to redis\n", key, value)
	return nil
}

type MemoryStore struct {
}

func (r MemoryStore) Store(key, value string) error {
	fmt.Printf("Storing (%s,%s) to file\n", key, value)
	return nil
}

func main() {
	var redisStore Store
	var memoryStore Store

	redisStore = RedisStore{}
	if err := redisStore.Store("001", "kosong kosong satu"); err != nil {
		fmt.Println("FAILED, redis")
	}

	memoryStore = MemoryStore{}
	if err := memoryStore.Store("002", "kosong kosong dua"); err != nil {
		fmt.Println("FAILED, memory")
	}
}
