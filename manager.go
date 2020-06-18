// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package go_cache

import (
	"github.com/just-coding-0/go-cache/internal/cacheMode"
	"github.com/just-coding-0/go-cache/internal/hashmap"
	"runtime"
)

type Manager struct {
	hashMap          hashmap.HashMap
	//set              Set
	cacheElimination cacheMode.CacheElimination
	Push             chan<- *cacheMode.Entry
	Pop              <-chan *cacheMode.Entry

	maxMemCap uint64

	mode uint8
}

var _ CacheManager = &Manager{}

type CacheManager interface {
	MapPut(key, value string, expiryPolicy int64) error // set key value
	MapLoad(key string) (string, bool)                  // get value
	MapRemove(key string)                               // remove

	//SetPut(setName, value string)
	//SetGetAll(setName string) []string
	//SetGetValueByLimit(setName string, limit uint16) []string
	//SetGetRandValue(setName string, limit uint16) []string

	Clear()
	Close()

	//flush() bool        // flush disk data
	//syncDiskData() bool // sync data on disk
}

//type Set interface {
//	Put(setName, value string)
//	GetAll(setName string) []string
//	GetValueByLimit(setName string, limit uint64) []string
//	GetRandValue(setName string, limit uint64) []string
//	Range(startKey string, endKey string) []string
//	Close()
//}

func NewManager(MaxMemCap uint64, shardSize uintptr, mode uint8) CacheManager {
	manager := &Manager{
		maxMemCap:        MaxMemCap,
		hashMap:          hashmap.NewHashMap(shardSize),
		mode:             mode,
		cacheElimination: NewCacheElimination(mode),
	}
	manager.Push = manager.cacheElimination.PushEntry()
	manager.Pop = manager.cacheElimination.PopEntry()

	// 监听chan
	go manager.WatchPopChan()

	return manager
}

func NewCacheElimination(mode uint8) cacheMode.CacheElimination {
	switch mode {
	case cacheMode.FIFO:
		return cacheMode.NewModeFIFO()
	//case cacheMode.LRU:
	//	return nil
	default:
		panic("未知的缓存策略")
	}
}

func (m *Manager) MapPut(key, value string, expiryPolicy int64) error {
	m.hashMap.Put(key, value)
	e := m.cacheElimination.GetEntry()
	err := e.SetExpiryTimes(expiryPolicy)
	if err != nil {
		m.cacheElimination.PutEntry(e)
		return err
	}
	e.SetKey(key)
	e.SetAction(cacheMode.Save)
	e.SetStoreType(cacheMode.HashMap)
	m.Push <- e
	return nil
}

func (m *Manager) MapLoad(key string) (value string, ok bool) {
	value, ok = m.hashMap.Load(key)

	if m.mode != cacheMode.FIFO {
		e := m.cacheElimination.GetEntry()
		e.SetKey(key)
		e.SetStoreType(cacheMode.HashMap)
		e.SetAction(cacheMode.Get)
		m.Push <- e
	}
	return
}

func (m *Manager) MapRemove(key string) {
	m.hashMap.Delete(key)
	e := m.cacheElimination.GetEntry()
	e.SetKey(key)
	e.SetStoreType(cacheMode.HashMap)
	e.SetAction(cacheMode.Delete)
	m.Push <- e
}

func (m *Manager) Clear() {
	// 重新初始化hashMap
	shardSize := m.hashMap.GetShardSize()
	m.hashMap = hashmap.NewHashMap(shardSize)

	runtime.GC()
}

func (m *Manager) Close() {
	m.cacheElimination.Close()
}

func (m *Manager) WatchPopChan() {
	for v := range m.Pop {
		if v.StoreType == cacheMode.HashMap {
			m.hashMap.Delete(v.Key)
			m.cacheElimination.PutEntry(v)
		}
	}

}
