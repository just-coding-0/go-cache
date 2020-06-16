// Copyright 2020 just-codeding-0 . All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cacheMode

import "time"

type ModeFIFO struct {
	cacheElimination
	BitMap []byte
	chain  *chain
}

const (
	bitSize = 8
	maxCnt  = 10000000
)

func NewModeFIFO() ModeFIFO {
	return ModeFIFO{cacheElimination: cacheElimination{
		Push:                 make(chan *Entry, 1000),
		Pop:                  make(chan *Entry, 10),
		refreshStatsExitChan: make(chan struct{}, 0), // 阻塞chan
		popExitChan:          make(chan struct{}, 0), // 阻塞Chan
		heapAlloc:            0,
		heapObjects:          0,
		MaxMem:               0,
	},
		BitMap: make([]byte, maxCnt), // 位表 10000000  10000016/1024/1024 = 9.5M
		chain:  newChain(),
	}
}

func (m *ModeFIFO) Start() {
	go m.refreshStats()
	go m.handlerPushMsg()
	go m.handlerPop()
}

func (m *ModeFIFO) handlerPushMsg() {

	for v := range m.Push {
		i := hash(v.Key)
		idx := i % maxCnt
		pos := i % bitSize
		if m.BitMap[idx]>>pos&1 == 1 {
			continue
		}
		m.BitMap[idx] |= 1 << pos
		m.chain.InsertFront(v)
	}
}

func (m *ModeFIFO) handlerPop() {
	var t = time.NewTimer(time.Second*10)

	for {
		select {
		case <-t.C:
			k := float64(m.heapAlloc) / float64(m.MaxMem)
			if k < 0.8 {
				break
			}
			cnt := 1000

			for i := 0; i < cnt; i++ {
				entry := m.chain.PopFromTail()
				i := hash(entry.Key)
				idx := i % maxCnt
				pos := i % bitSize
				m.BitMap[idx] ^= 1 << pos

				if entry == nil {
					break
				}
				m.Pop <- entry
			}
		case <-m.popExitChan:
			t.Stop()
			close(m.popExitChan)
		}
	}
}
