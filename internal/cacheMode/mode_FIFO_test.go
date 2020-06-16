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
	f.SetMaxMemUsed(13 * MB) // 50MB
	f.Start()
	c := f.PushEnter()
	c1 := f.PopEnter()
	fmt.Println(time.Now())
	go func() {
		for {
			c <- &Entry{Key: fmt.Sprintf("%d", rand.Int()), ExpiryPolicy: time.Now().Unix()}
			time.Sleep(time.Microsecond * 200)
		}
	}()

	var cnt int
	for v := range c1 {
		cnt++
		t := time.Unix(v.ExpiryPolicy, 0)
		fmt.Println(t)
		if cnt == 10 {
			break
		}
	}
}
