// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

import (
	"fmt"
	"hash/fnv"
	"runtime"
	"sync"
	"time"
)

const (
	FIFO = iota // 0 First In First Out
	LRU         // 1 Least Recently Used
)

/*
FIFO：First In First Out，先进先出。判断被存储的时间，离目前最远的数据优先被淘汰。
LRU：Least Recently Used，最近最少使用。判断最近被使用的时间，目前最远的数据优先被淘汰。
*/

const (
	B  = 1
	Kb = B * 1024
	MB = Kb * 1024
	GB = MB * 1024
)

/*
该接口会实现两个功能
1.往里面放entry
2.返回一个订阅的chan,该类会将需要淘汰的entry返回

<- f
f <-
*/
type CacheElimination interface {
	PushEntry() chan<- *Entry
	PopEntry() <-chan *Entry
	SetMaxMemUsed(maxMem uintptr)
	Close()
	Start()

	GetEntry() *Entry
	PutEntry(e *Entry)
}

type cacheElimination struct {
	push                 chan *Entry
	pop                  chan *Entry
	pool                 sync.Pool
	popExitChan          chan struct{}
	refreshStatsExitChan chan struct{}
	heapAlloc            uint64
	heapObjects          uint64
	MaxMem               uintptr // byte
	debug                bool
}

func (c *cacheElimination) PushEntry() chan<- *Entry {
	if c.push != nil {
		return c.push
	}
	panic("not found push chan")
	return nil
}

func (c *cacheElimination) PopEntry() <-chan *Entry {
	if c.pop != nil {
		return c.pop
	}
	panic("not found pop chan")
	return nil
}

func (c *cacheElimination) refreshStats() {

	stats := &runtime.MemStats{}

	t := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(stats)
			c.heapAlloc = stats.HeapAlloc
			c.heapObjects = stats.HeapObjects
			if c.debug {
				fmt.Println(c.heapAlloc / 1024 / 1024)
			}
		case <-c.refreshStatsExitChan:
			t.Stop()
			close(c.refreshStatsExitChan)
			return
		}
	}

}

func (c *cacheElimination) Close() {
	c.refreshStatsExitChan <- struct{}{}
	c.popExitChan <- struct{}{}
	close(c.push)
	close(c.pop)
}

func (c *cacheElimination) SetMaxMemUsed(maxMem uintptr) {
	c.MaxMem = maxMem
}

func (c *cacheElimination) GetEntry() *Entry {
	e := c.pool.Get().(*Entry)
	e.reset()
	return e
}
func (c *cacheElimination) PutEntry(e *Entry) {
	c.pool.Put(e)
}

func hash(v interface{}) uint64 {
	switch v.(type) {
	case string:
		h := fnv.New64a()
		h.Write([]byte(v.(string)))
		return h.Sum64()
	default:
		panic("the type not supported!")
	}
}
