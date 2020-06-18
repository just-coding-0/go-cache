// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestEntry_SetAction(t *testing.T) {
	n := newEntry()
	n.SetAction(Save)
	assert.Equal(t, Save,n.Action)
}

func TestEntry_SetKey(t *testing.T) {
	n := newEntry()
	key:="hello"
	n.SetKey(key)
	assert.Equal(t, key,n.Key)
}

func TestEntry_SetExpiryTimes(t *testing.T) {
	n := newEntry()
	unix := time.Now().AddDate(0, 0, 1).Unix()
	n.SetExpiryTimes(unix)
	assert.Equal(t, unix,n.ExpiryTimes)
}

func TestEntry_SetStoreType(t *testing.T) {
	n := newEntry()
	n.SetStoreType(HashMap)
	assert.Equal(t, n.Action, HashMap)

}
