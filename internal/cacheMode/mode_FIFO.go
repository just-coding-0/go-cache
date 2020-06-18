// Copyright 2020 just-codeding-0 . All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cacheMode

import (
	"container/heap"
	"sync"
	"time"
)

type ModeFIFO struct {
	cacheElimination

	chainM    sync.Map
	chainLock sync.Mutex
	chain     *chain

	heapM    sync.Map
	heapLock sync.Mutex
	heap     *entryHeap
}

func NewModeFIFO() *ModeFIFO {
	m := &ModeFIFO{cacheElimination: cacheElimination{
		push:                 make(chan *Entry, 10000),
		pop:                  make(chan *Entry, 100),
		refreshStatsExitChan: make(chan struct{}, 0), // 阻塞chan
		popExitChan:          make(chan struct{}, 0), // 阻塞Chan
		heapAlloc:            0,
		heapObjects:          0,
		MaxMem:               0,
	},
		chainM: sync.Map{},
		chain:  newChain(),

		heap:  newEntryHeap(10000), // 默认设置一万 80000byte = 78kb
		heapM: sync.Map{},
	}

	m.pool = sync.Pool{}
	m.pool.New = func() interface{} {
		return newEntry()
	}

	return m
}

func (m *ModeFIFO) Start() {
	go m.refreshStats()
	go m.handlerPushMsg()
	go m.handlerPop()
}

func (m *ModeFIFO) handlerPushMsg() {

	for e := range m.push {

		switch e.Action {
		case Save:
			m.handlerSave(e)
		case Delete:
			m.handlerDelete(e.Key)
		}
	}
}

func (m *ModeFIFO) handlerDelete(key string) {

	if val, ok := m.heapM.Load(key); ok {

		m.heapLock.Lock()
		m.heap.deleteEntry(val.(*Entry))
		m.heapLock.Unlock()

		m.PutEntry(val.(*Entry))
		m.heapM.Delete(key)

	}

	if node, ok := m.chainM.Load(key); ok {

		m.chainLock.Lock()
		entry := m.chain.Pop(node.(*Node))
		m.chainLock.Unlock()

		m.PutEntry(entry)
		m.chainM.Delete(key)
	}
}

func (m *ModeFIFO) handlerSave(v *Entry) {
	if v.ExpiryTimes > 0 { // 使用堆管理

		if val, ok := m.heapM.Load(v.Key); ok { // 如果已经存在,就将该entry放回缓存池
			if val.(*Entry).ExpiryTimes != v.ExpiryTimes {
				val.(*Entry).ExpiryTimes = v.ExpiryTimes
				m.heapLock.Lock()
				m.heap.changeEntry(val.(*Entry))
				m.heapLock.Unlock()
			}

			m.PutEntry(v)
			return
		}

		m.heapM.Store(v.Key, v)

		m.heapLock.Lock()
		heap.Push(m.heap, v)
		m.heapLock.Unlock()

	} else {
		if val, ok := m.chainM.Load(v.Key); ok { //  如果已经存在,就将该entry放回缓存池
			m.chainLock.Lock()
			m.chain.MoveNodeToFront(val.(*Node))
			m.chainLock.Unlock()

			m.PutEntry(v)
			return
		}

		n := m.chain.getNewNode()
		n.setValue(v)
		m.chainM.Store(v.Key, n)

		m.chainLock.Lock()
		m.chain.InsertFront(n)
		m.chainLock.Unlock()
	}
}

func (m *ModeFIFO) handlerPop() {
	var t = time.NewTicker(time.Second * 1)
	var unix int64
	for {
		select {
		case <-t.C:
			unix = time.Now().Unix()
			m.heapLock.Lock()
			for m.heap.Len() > 0 && m.heap.items[m.heap.Len()-1].ExpiryTimes <= unix { // 最小堆,会将即将过期的key,传输出去
				e := m.heap.Pop().(*Entry)
				m.heapM.Delete(e.Key)
				m.pop <- e
			}
			m.heapLock.Unlock()

			k := float64(m.heapAlloc) / float64(m.MaxMem)
			if k < 0.8 {
				break
			}
			cnt := 1000

			for i := 0; i < cnt; i++ {
				m.chainLock.Lock()
				entry := m.chain.PopFromTail()
				m.chainLock.Unlock()
				if entry == nil {
					break
				}

				m.chainM.Delete(entry.Key)
				m.pop <- entry
			}
		case <-m.popExitChan:
			t.Stop()
			close(m.popExitChan)
		}
	}
}
