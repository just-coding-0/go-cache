// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/just-coding-0/go-cache/cmd/manager/handler/hashmap"
	"github.com/just-coding-0/go-cache/cmd/manager/middleware"
)

func hashMapRoute(e *gin.Engine) {
	g := e.Group("/map/", middleware.CheckIp())
	g.GET("/load", hashmap.LoadHandler)
	g.GET("/remove", hashmap.RemoveHandler)
	g.GET("/put", hashmap.PutHandler)
}
