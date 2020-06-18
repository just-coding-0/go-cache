// Copyright 2020 just-codeding-0 . All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cacheMode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestModeFIFO(t *testing.T) {
	f := NewModeFIFO()
	f.debug = true
	f.SetMaxMemUsed(50 * MB) // 50MB
	f.Start()
	c := f.PushEntry()
	c1 := f.PopEntry()
	go func() {
		for {
			c <- &Entry{key: fmt.Sprintf("%d", rand.Int()), expiryTimes: time.Now().Unix()}
			time.Sleep(time.Microsecond * 200)
		}
	}()

	go func() {
		for {
			c <- &Entry{key: fmt.Sprintf("%d", rand.Int()), expiryTimes: time.Now().Unix()}
			time.Sleep(time.Microsecond * 200)
		}
	}()

	go func() {
		for {
			c <- &Entry{key: fmt.Sprintf("%d", rand.Int()), expiryTimes: time.Now().Unix()}
			time.Sleep(time.Microsecond * 200)
		}
	}()

	var cnt int
	for  range c1 {
		cnt++
		if cnt == 100000 {
			break
		}
	}
}
