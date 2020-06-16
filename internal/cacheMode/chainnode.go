// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

import "sync"

type chain struct {
	Front *Node
	Tail  *Node
	pool  sync.Pool
}

type Node struct {
	value *Entry
	Pre   *Node
	Next  *Node
}

func newChain() *chain {
	c := &chain{pool: sync.Pool{}}
	c.pool.New = func() interface{} {
		return newNode()
	}

	return c
}

func newNode() *Node {
	return &Node{}
}

func (n *Node) setValue(entry *Entry) {
	n.value = entry
}

func (n *Node) reset() {
	n.Pre = nil
	n.Next = nil
	n.value = nil
}

func (c *chain) InsertFront(entry *Entry) {
	n := c.pool.Get().(*Node)
	n.setValue(entry)

	if c.Front == nil {
		c.Front = n
		c.Tail = n
		return
	}

	c.Front.Pre = n
	n.Next = c.Front
	c.Front = n
}

func (c *chain) InsertTail(entry *Entry) {
	n := c.pool.Get().(*Node)
	n.setValue(entry)

	if c.Tail == nil {
		c.Front = n
		c.Tail = n
		return
	}

	n.Pre = c.Tail
	c.Tail.Next = n
	c.Tail = n
}

func (c *chain) PopFromTail() *Entry {
	if c.Tail == nil {
		return nil
	}

	if c.Tail == c.Front {
		value := c.Tail.value
		c.Tail.reset()
		c.pool.Put(c.Tail)
		c.Tail = nil
		c.Front = nil

		return value
	}

	tail := c.Tail
	c.Tail = c.Tail.Pre
	if c.Tail != nil {
		c.Tail.Next = nil
	}

	value := tail.value
	tail.reset()
	c.pool.Put(tail)

	return value
}

func (c *chain) PopFromFront() *Entry {
	if c.Front == nil {
		return nil
	}

	if c.Front == c.Tail {
		value := c.Front.value
		c.Front.reset()
		c.pool.Put(c.Front)
		c.Tail = nil
		c.Front = nil

		return value
	}

	front := c.Front
	c.Front = front.Next
	if c.Front != nil {
		c.Front.Pre = nil
	}

	value := front.value

	front.reset()
	c.pool.Put(front)
	return value
}

func (c *chain) Pop(n *Node) *Entry {

	value := n.value

	pre := n.Pre
	next := n.Next

	if n.Pre != nil {
		pre.Next = next
		if next != nil {
			next.Pre = pre
		}
	}

	n.reset()
	c.pool.Put(n)

	return value
}
