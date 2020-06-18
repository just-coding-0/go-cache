// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

import (
	"container/heap"
	"unsafe"
)

type entryHeap struct {
	items []*Entry
	m     map[uintptr]int
}

var _ heap.Interface = &entryHeap{}

func newEntryHeap(size uint16) *entryHeap {
	var entryHeap = &entryHeap{}
	entryHeap.items = make([]*Entry, 0, size)
	entryHeap.m = make(map[uintptr]int)
	return entryHeap
}
func (h *entryHeap) Len() int {
	return len(h.items)
}

func (h *entryHeap) Less(i, j int) bool {
	return h.items[i].ExpiryTimes < h.items[j].ExpiryTimes
}

// Swap swaps the elements with indexes i and j.
func (h *entryHeap) Swap(i, j int) {
	h.m[uintptr(unsafe.Pointer(h.items[i]))] = j
	h.m[uintptr(unsafe.Pointer(h.items[j]))] = i
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *entryHeap) Push(x interface{}) {
	e := x.(*Entry)
	h.m[uintptr(unsafe.Pointer(e))] = h.Len()
	h.items = append(h.items, e)
}

func (h *entryHeap) Pop() interface{} {
	l := h.Len()
	item := h.items[l-1]
	h.items = h.items[:l-1]
	delete(h.m, uintptr(unsafe.Pointer(item)))
	return item
}

func (h *entryHeap) deleteEntry(e *Entry) {
	idx := h.m[uintptr(unsafe.Pointer(e))]
	heap.Remove(h, idx)
}

func (h *entryHeap) changeEntry(e *Entry) {
	idx := h.m[uintptr(unsafe.Pointer(e))]
	heap.Fix(h, idx)
}
