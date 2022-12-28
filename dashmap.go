package dashmap

import "sync"

type Entry[K comparable, V any] struct {
	Key K
	Value V
}

type SyncMapper[K comparable, V any] interface {
	Get(key K) (V, bool)
	Put(key K, value V)
	Entries() []Entry[K, V]
}

type DashMap[K comparable, V any] struct {
	lock sync.RWMutex
	mp map[K]V
}

var _ SyncMapper[int, any] = (*DashMap[int, any])(nil)

func New[K comparable, V any]() *DashMap[K, V] {
	return &DashMap[K, V]{
		mp: make(map[K]V),
	}
}

func (dm *DashMap[K, V]) Get(key K) (V, bool) {
	dm.lock.RLock()
	defer dm.lock.RUnlock()
	val, ok := dm.mp[key]
	return val, ok
}

func (dm *DashMap[K, V]) Put(key K, value V) {
	dm.lock.Lock()
	defer dm.lock.Unlock()
	dm.mp[key] = value
}

func (dm *DashMap[K, V]) Entries() []Entry[K, V] {
	dm.lock.RLock()
	defer dm.lock.RUnlock()
	entries := make([]Entry[K, V], len(dm.mp))
	idx := 0
	for k, v := range dm.mp {
		entries[idx] = Entry[K, V]{Key: k, Value: v}
		idx++
	}
	return entries
}
