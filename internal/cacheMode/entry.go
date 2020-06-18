// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

import (
	"errors"
	"time"
)

const MaxExpiryPolicy = 60 * 60 * 24 * 30

var (
	InvalidExpiryTimes = errors.New("无效的到期时间")
)

type storeType uint8

const (
	HashMap storeType = 0
	Set     storeType = 1
)

type ActionType uint8

const (
	Save   ActionType = 0
	Get    ActionType = 1
	Delete ActionType = 2
)

type Entry struct {
	Key         string // key
	ExpiryTimes int64  //  Timestamp
	StoreType   storeType
	Action      ActionType
}

func (e *Entry) reset() {
	e.StoreType = 0
	e.Key = ""
	e.ExpiryTimes = 0
	e.Action = 0
}

func (e *Entry) SetKey(key string) {
	e.Key = key
}

func (e *Entry) SetStoreType(storeType storeType) {
	e.StoreType = storeType
}

func (e *Entry) SetAction(action ActionType) {
	e.Action = action
}

func (e *Entry) SetExpiryTimes(expiryTimes int64) error {
	n := time.Now().Unix()
	if expiryTimes < n {
		return InvalidExpiryTimes
	}

	// 最长一周
	if expiryTimes > n+MaxExpiryPolicy {
		return InvalidExpiryTimes
	}

	e.ExpiryTimes = expiryTimes
	return nil
}


func newEntry() *Entry {
	return &Entry{}
}
