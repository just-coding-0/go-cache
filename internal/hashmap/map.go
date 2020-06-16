// Copyright 2020 just-codeding-0 . All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package hashmap

import (
	"math"
	"reflect"
	"sync"
	"unsafe"
)

type hashMap struct {
	shards     []shard
	hash0      uintptr
	hasher     func(unsafe.Pointer, uintptr) uintptr
	shardCount uintptr
}

type shard struct {
	syncMap sync.Map
	//Map map[string]string
	//lock sync.Mutex
}

const (
	shardSize = unsafe.Sizeof(shard{})
)

// 最小255个分片
func NewHashMap(shardCount uintptr) *hashMap {
	//if shardCount < math.MaxUint8 {
	//	shardCount = math.MaxUint8
	//}
	if shardCount > math.MaxUint16 {
		shardCount = math.MaxUint16
	}
	var shards = make([]shard, shardCount)

	for idx := range shards {
		shards[idx] = shard{
			syncMap: sync.Map{},
		}
	}
	h := &hashMap{}
	h.shards = shards
	h.shardCount = shardCount
	h.init()

	return h
}

// 使用map的hash0
func (h *hashMap) init() {
	k := make(map[string]struct{})
	value := reflect.ValueOf(k)
	k1 := *(**hmap)(unsafe.Pointer(&k))
	tt := *(**mapType)(unsafe.Pointer(&value))
	h.hash0 = uintptr(k1.hash0)
	h.hasher = tt.hasher
}

func (h *hashMap) Delete(key string) bool {
	return h.delete(key)
}

func (h *hashMap) delete(key string) bool {
	shard := (*shard)(add(unsafe.Pointer(&h.shards[0]), h.getShardNumber(key)*shardSize))
	shard.syncMap.Delete(key)
	return true
}

func (h *hashMap) Put(key, value string) {
	h.put(key, value)
}

func (h *hashMap) put(key, value string) {
	shard := (*shard)(add(unsafe.Pointer(&h.shards[0]), h.getShardNumber(key)*shardSize))
	shard.syncMap.Store(key, value)
}

func (h *hashMap) Load(key string) (string, bool) {
	return h.load(key)
}

func (h *hashMap) load(key string) (string, bool) {
	shard := (*shard)(add(unsafe.Pointer(&h.shards[0]), h.getShardNumber(key)*shardSize))
	val, ok := shard.syncMap.Load(key)
	if !ok {
		return "", false
	}
	return val.(string), true
}

func (h *hashMap) GetShardSize() uintptr {
	return h.shardCount
}

func (h *hashMap) getShardNumber(key string) uintptr {
	hash := h.hasher(unsafe.Pointer(&key), h.hash0)
	return hash % h.shardCount
}

func (h *hashMap) Close(){
	h.shards = nil
	h.hasher = nil
}

// 使用指针操作,可以避免边界检查
func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}
