// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/just-coding-0/go-cache/cmd/manager/handler"
	"net/http"
	"sync"
)

var ipMap *sync.Map

func SetIpMap(_ipMap *sync.Map) {
	ipMap = _ipMap
}

func CheckIp() func(c *gin.Context) {
	return func(c *gin.Context) {
		_, ok := ipMap.Load(c.ClientIP())
		if !ok  {
			c.JSON(http.StatusUnauthorized, handler.ApiResponse{
				Msg: "该ip未验证",
			})
			c.Abort()
		}
		c.Next()
	}

}
