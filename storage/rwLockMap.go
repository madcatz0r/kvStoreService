package storage

import (
	"fmt"
	"sync"
)

type RWLockMap struct {
	sync.RWMutex
	internal map[string]interface{}
}

func (m *RWLockMap) Put(key string, value []byte) {
	m.Lock()
	m.internal[key] = value
	m.Unlock()
}

func (m *RWLockMap) Get(key string) ([]byte, error) {
	m.RLock()
	res, ok := m.internal[key]
	m.RUnlock()
	if !ok {
		return []byte{}, fmt.Errorf("key %v not found", key)
	}
	val, ok := res.([]byte)
	if !ok {
		return []byte{}, fmt.Errorf("unexpected cast error, key: %v, value: %+v", key, res)
	}
	return val, nil
}

func (m *RWLockMap) Delete(key string) error {
	_, err := m.Get(key)
	if err != nil {
		return err
	}
	m.Lock()
	delete(m.internal, key)
	m.Unlock()
	return nil
}
