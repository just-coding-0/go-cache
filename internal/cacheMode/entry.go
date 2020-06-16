// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cacheMode

const DefaultExpiryPolicy = 60 * 60
const MaxExpiryPolicy = 60 * 60 * 24

type storeType uint8

const (
	hashMap storeType = 1
	set     storeType = 1
)

type Entry struct {
	storeType    storeType
	Key          string // key
	ExpiryPolicy int64  //  Timestamp
}

func NewEntry() *Entry {
	return &Entry{}
}
