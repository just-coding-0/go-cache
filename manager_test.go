package go_cache

import (
	"fmt"
	"github.com/just-coding-0/go-cache/internal/cacheMode"
	"testing"
	"time"
)

func TestNewManager(t *testing.T) {
	manager := NewManager(1024*1024*1024, 256, cacheMode.FIFO)
	err := manager.MapPut("hello", "val", time.Now().Unix())
	fmt.Println(err)
	//

	time.Sleep(time.Millisecond*100)
	fmt.Println(manager.MapLoad("hello"))

}
