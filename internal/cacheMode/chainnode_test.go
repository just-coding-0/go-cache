// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChain_InsertFront(t *testing.T) {
	c := newChain()
	var tests = []string{"1", "2", "3", "4", "5"}

	for _, v := range tests {
		c.InsertFront(&Entry{
			Key: v,
		})
	}

	println("倒序")

	for _, v := range tests {
		entry := c.PopFromTail()
		fmt.Println(entry.Key, v)
		assert.Equal(t, entry.Key, v)
	}

	for _, v := range tests {
		c.InsertFront(&Entry{
			Key: v,
		})
	}

	println("倒序")

	for i := len(tests) - 1; i >= 0; i-- {
		entry := c.PopFromFront()
		fmt.Println(entry.Key, tests[i])

		assert.Equal(t, entry.Key, tests[i])
	}

}

func TestChain_InsertTail(t *testing.T) {
	c := newChain()
	var tests = []string{"1", "2", "3", "4", "5"}

	for _, v := range tests {
		c.InsertTail(&Entry{
			Key: v,
		})
	}

	for _, v := range tests {
		entry := c.PopFromFront()
		assert.Equal(t, v, entry.Key)
	}
}

func TestChain_PopFromFront(t *testing.T) {
	c := newChain()
	var tests = []string{"1", "2", "3", "4", "5"}

	for _, v := range tests {
		c.InsertTail(&Entry{
			Key: v,
		})
	}

	for _, v := range tests {
		entry := c.PopFromFront()
		assert.Equal(t, v, entry.Key)
	}
}

func TestChain_PopFromTail(t *testing.T) {
	c := newChain()
	var tests = []string{"1", "2", "3", "4", "5"}

	for _, v := range tests {
		c.InsertTail(&Entry{
			Key: v,
		})
	}

	for i := len(tests) - 1; i >= 0; i-- {
		entry := c.PopFromTail()
		assert.Equal(t, tests[i], entry.Key)
	}
}
