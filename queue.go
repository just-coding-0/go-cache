// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package go_cache

import (
	"math"
	"sync"
)

const (
	maxSameSizeGrow = math.MaxUint16
)

type RoundQueue interface {
	Insert(uint32)
	Pop() uint32
	Growing()
	IsEmpty() bool
}

type roundQueue struct {
	Queue []uint32 // 队列
	Count uint64   // 当前队列内可用元素数量
	Cap   uint64   // 长度
	Front uint64   // 头部
	Rear  uint64   // 尾部
	Lock  sync.Mutex
}

func NewRoundQueue(cap uint64) *roundQueue {
	if cap < 100 {
		cap = 100
	}

	if cap > maxSameSizeGrow {
		cap = maxSameSizeGrow
	}

	return &roundQueue{
		Queue: make([]uint32, cap),
		Count: 0,
		Cap:   cap,
		Front: 0,
		Rear:  0,
		Lock:  sync.Mutex{},
	}
}
func (r *roundQueue) Insert(index uint32) {
	r.Growing()
	r.Lock.Lock()
	r.Queue[r.Rear] = index
	r.Rear++
	r.Count++
	if r.Rear == r.Count {
		r.Rear = 0
	}
	r.Lock.Unlock()
}

func (r *roundQueue) Pop() (index uint32) {
	if r.IsEmpty() {
		return
	}
	r.Lock.Lock()
	r.Count--
	index = r.Queue[r.Front]
	r.Front++
	if r.Front == r.Count { // 重置头部
		r.Front = 0
	}
	r.Lock.Unlock()
	return
}

// 扩容
func (r *roundQueue) Growing() {
	// 只有当前已用空间接近已分配空间的0.9以上,则扩容 first path
	if r.Count < uint64(float64(r.Cap)*0.9) {
		return
	}
	r.Lock.Lock()
	if r.Count < uint64(float64(r.Cap)*0.9) { // 防止多次调用
		return
	}

	var newCap uint64
	if r.Cap > maxSameSizeGrow {
		newCap = r.Cap + (r.Cap >> 2) // new_size =  cap + cap/4
	} else {
		newCap = r.Cap << 1 // new_size = r.cap * 2
	}
	newQueue := make([]uint32, newCap, newCap)
	copy(newQueue, r.Queue)
	r.Queue = nil
	r.Queue = newQueue
	r.Cap = newCap
	r.Lock.Unlock()
}

func (r *roundQueue) IsEmpty() bool {
	return r.Count == 0
}
