// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package go_cache

import (
	"math/rand"
	"testing"
)


func BenchmarkRoundQueue(b *testing.B) {
	queue := NewRoundQueue(100)

	for i := 0; i < b.N; i++ {
		queue.Insert(rand.Uint32())
	}

	for i := 0; i < b.N; i++ {
		queue.Pop()
	}

	b.Log(queue.Count,queue.Rear,queue.Front,queue.Cap)

}
