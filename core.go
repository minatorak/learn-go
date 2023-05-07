package main

import (
	"errors"
	"sync"
)

var syncMap = struct {
	sync.RWMutex
	store map[string]string
}{store: make(map[string]string)}

var ErroNoSuchKey = errors.New("no such key")

func Put(key string, value string) error {
	syncMap.Lock()
	syncMap.store[key] = value
	syncMap.Unlock()
	return nil
}

func Get(key string) (string, error) {
	syncMap.RLock()
	value, ok := syncMap.store[key]
	syncMap.RUnlock()
	if !ok {
		return "", ErroNoSuchKey
	}
	return value, nil
}

func Delete(key string) error {
	syncMap.Lock()
	delete(syncMap.store, key)
	syncMap.Unlock()
	return nil
}

func DetactError(e error) {
	if errors.Is(e, ErroNoSuchKey) {
		println("e is erroNoSuchKey")
	}
}
