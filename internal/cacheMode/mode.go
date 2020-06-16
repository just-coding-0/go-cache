// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

import (
	"fmt"
	"hash/fnv"
	"runtime"
	"time"
)

const (
	LRU  = iota // 0 Least Recently Used
	LFU         // 1 least frequently used
	FIFO        // 2 First In First Out
)

const (
	B  = 1
	Kb = B * 1024
	MB = Kb * 1024
	GB = MB * 1024
)

/*
该接口会实现两个功能
1.往里面放entry
2.返回一个订阅的chan,该类会将需要淘汰的enter返回

<- f
f <-
*/
type CacheElimination interface {
	PushEnter() chan *Entry
	PopEnter() chan *Entry
	SetMaxMemUsed(maxMem uintptr)
	Close()
	Start()
}

type cacheElimination struct {
	Push                 chan *Entry
	Pop                  chan *Entry
	popExitChan          chan struct{}
	refreshStatsExitChan chan struct{}
	heapAlloc            uint64
	heapObjects          uint64
	MaxMem               uintptr // byte
	debug                bool
}

func (c *cacheElimination) PushEnter() chan<- *Entry {
	if c.Push != nil {
		return c.Push
	}
	panic("not found push chan")
	return nil
}

func (c *cacheElimination) PopEnter() <-chan *Entry {
	if c.Pop != nil {
		return c.Pop
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
	close(c.Push)
	close(c.Pop)
}

func (c *cacheElimination) SetMaxMemUsed(maxMem uintptr) {
	c.MaxMem = maxMem
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
