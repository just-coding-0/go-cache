package go_cache1

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
