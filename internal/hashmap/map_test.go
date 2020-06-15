package hashmap_test

import (
	"fmt"
	"github.com/just-coding-0/go-cache/internal/hashmap"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func BenchmarkNewHashMap256(b *testing.B) {
	rand.Seed(time.Now().Unix())
	var wait sync.WaitGroup

	b.StartTimer()
	m := hashmap.NewHashMap(255)

	for i := 0; i < 100; i++ {
		wait.Add(1)
		go func() {
			defer func() {
				wait.Done()
			}()
			for i := 0; i < 1000; i++ {
				key := fmt.Sprintf("%d ", rand.Int())
				m.Put(key,key)
				m.Load(key)
			}
		}()
	}
	wait.Wait()

}

func TestHashMap_Delete(t *testing.T) {
	m := hashmap.NewHashMap(256)
	phones := map[string]string{
		"apple":  "9plus",
		"vivo":   "x50 pro",
		"HUAWEI": "mate Xs",
		"MI":     "MIX3",
	}

	for k, v := range phones {
		m.Put(k, v)
	}

	for k := range phones {
		m.Delete(k)
	}

	for k := range phones {
		k, ok := m.Load(k)
		assert.Empty(t, k)
		assert.True(t, !ok)

	}
}

func TestHashMap_Load(t *testing.T) {
	m := hashmap.NewHashMap(256)
	phones := map[string]string{
		"apple":  "9plus",
		"vivo":   "x50 pro",
		"HUAWEI": "mate Xs",
		"MI":     "MIX3",
	}

	for k, v := range phones {
		m.Put(k, v)
	}

	for k, v := range phones {
		val, ok := m.Load(k)
		assert.True(t, ok)
		assert.Equal(t, val, v)
	}
}

func TestHashMap_Put(t *testing.T) {
	m := hashmap.NewHashMap(256)
	phones := map[string]string{
		"apple":  "9plus",
		"vivo":   "x50 pro",
		"HUAWEI": "mate Xs",
		"MI":     "MIX3",
	}

	for k, v := range phones {
		m.Put(k, v)
	}
}

func TestHashMap_GetShardSize(t *testing.T) {
	m := hashmap.NewHashMap(256)

	assert.Equal(t, m.GetShardSize(), uintptr(256))
}
