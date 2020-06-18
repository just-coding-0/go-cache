// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestChain_InsertFront(t *testing.T) {
	c := newChain()
	var tests = []string{"1", "2", "3", "4", "5"}

	for _, v := range tests {
		n := c.getNewNode()
		n.value = &Entry{
			key: v,
		}
		c.InsertFront(n)
	}

	println("倒序")

	for _, v := range tests {
		entry := c.PopFromTail()
		fmt.Println(entry.key, v)
		assert.Equal(t, entry.key, v)
	}

	for _, v := range tests {
		n := c.getNewNode()
		n.value = &Entry{
			key: v,
		}
		c.InsertFront(n)
	}

	println("倒序")

	for i := len(tests) - 1; i >= 0; i-- {
		entry := c.PopFromFront()
		fmt.Println(entry.key, tests[i])

		assert.Equal(t, entry.key, tests[i])
	}

}

func TestChain_InsertTail(t *testing.T) {
	c := newChain()
	var tests = []string{"1", "2", "3", "4", "5"}

	for _, v := range tests {
		n := c.getNewNode()
		n.value = &Entry{
			key: v,
		}
		c.InsertTail(n)
	}

	for _, v := range tests {
		entry := c.PopFromFront()
		assert.Equal(t, v, entry.key)
	}
}

func TestChain_PopFromFront(t *testing.T) {
	c := newChain()
	var tests = []string{"1", "2", "3", "4", "5"}

	for _, v := range tests {
		n := c.getNewNode()
		n.value = &Entry{
			key: v,
		}
		c.InsertTail(n)
	}

	for _, v := range tests {
		entry := c.PopFromFront()
		assert.Equal(t, v, entry.key)
	}
}

func TestChain_PopFromTail(t *testing.T) {
	c := newChain()
	var tests = []string{"1", "2", "3", "4", "5"}

	for _, v := range tests {
		n := c.getNewNode()
		n.value = &Entry{
			key: v,
		}
		c.InsertTail(n)
	}

	for i := len(tests) - 1; i >= 0; i-- {
		entry := c.PopFromTail()
		assert.Equal(t, tests[i], entry.key)
	}
}

func TestChain_SycPool(t *testing.T) {

	c := newChain()

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	for j := 0; j < 1000000; j++ {

		for i := 0; i < 10; i++ {
			n := c.getNewNode()
			n.value = newEntry()
			c.InsertFront(n)
		}

		for {
			node := c.PopFromFront()
			if node == nil {
				break
			}

		}
	}

	fmt.Println(stats.HeapAlloc/1024, "kb    ", stats.HeapObjects)
	runtime.ReadMemStats(&stats)
	fmt.Println("before")
	fmt.Println(stats.HeapAlloc/1024, "kb    ", stats.HeapObjects)

}
