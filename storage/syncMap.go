package storage

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	sync.Map
}

func (s *SyncMap) Put(key string, value []byte) {
	s.Store(key, value)
	return
}

func (s *SyncMap) Get(key string) ([]byte, error) {
	value, ok := s.Load(key)
	if !ok {
		return nil, fmt.Errorf("key %v not found", key)
	}
	val, ok := value.([]byte)
	if !ok {
		return nil, fmt.Errorf("unexpected cast error, key: %v, value: %#v", key, value)
	}
	return val, nil
}

func (s *SyncMap) Delete(key string) error {
	_, ok := s.Load(key)
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}
	s.Map.Delete(key)
	return nil
}
