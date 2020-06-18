// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/just-coding-0/go-cache/cmd/manager/handler"
	"github.com/just-coding-0/go-cache/cmd/manager/middleware"
	"github.com/just-coding-0/go-cache/cmd/manager/router"
	"github.com/just-coding-0/go-cache/cmd/manager/utils"
	"runtime"
	"sync"
	"time"
)

var (
	mode       = flag.Uint("mode", 0, "set cache mode [ 0:FIFO , 1:LRU ]")
	port       = flag.Int("port", 10001, "set port")
	privateKey = flag.String("privatekey", "/Users/ted/go/src/github.com/just-coding-0/go-cache/rsa_private_key.pem", " rsa private key")
	maxMem     = flag.Uint64("max_mem", 1024*1024*1024, "max use mem ")
	shardSize  = flag.Uint64("share_size", 256, "shard_size ")
	ipMap      = sync.Map{}
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	err := handler.SetPrivateKey(*privateKey)
	if err != nil {
		panic(err)
	}

	closeManager := handler.SetCacheManager(*maxMem, uintptr(*shardSize), uint8(*mode))

	defer func() {
		closeManager()
	}()

	handler.SetIpMap(&ipMap)
	middleware.SetIpMap(&ipMap)
	router.IndexRouter(r)

	// ticker
	ticker := utils.NewTicker(func() {
		date := time.Now().Format(`2006-01-02`)
		ipMap.Range(func(key, value interface{}) bool {
			if value != date {
				ipMap.Delete(key)
			}
			return false
		})
	}, 6)

	go ticker.Start()
	defer ticker.Stop()

	err = r.Run(fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println(err)
	}
}
