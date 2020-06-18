// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package hashmap

import (
	"github.com/gin-gonic/gin"
	. "github.com/just-coding-0/go-cache/cmd/manager/handler"
	v1 "github.com/just-coding-0/go-cache/v1"
	"net/http"
)

func LoadHandler(c *gin.Context) {
	data := c.Query("request")
	request, err := v1.DecodeLoadRequest(data)
	if CheckErrWithCode(err, DecodeError, c) {
		return
	}

	value, ok := Manager.MapLoad(request.Key)

	res := &v1.LoadResponse{}
	res.Value = value
	res.Ok = ok

	str, err := v1.EncodeLoadResponse(res)
	if CheckErrWithCode(err, EncodeError, c) {
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Data: str,
		Code: 0,
	})
}

func RemoveHandler(c *gin.Context) {
	data := c.Query("request")
	request, err := v1.DecodeRemoveRequest(data)
	if CheckErrWithCode(err, DecodeError, c) {
		return
	}

	Manager.MapRemove(request.Key)

	res := &v1.RemoveResponse{}
	res.Ok = true

	data, err = v1.EncodeRemoveResponse(res)
	if CheckErrWithCode(err, EncodeError, c) {
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Data: data,
		Code: 0,
	})
}

func PutHandler(c *gin.Context) {
	data := c.Query("request")
	request, err := v1.DecodePutRequest(data)
	if CheckErrWithCode(err, DecodeError, c) {
		return
	}

	err = Manager.MapPut(request.Key, request.Value, request.ExpiryPolicy)
	if CheckErrWithCode(err, MapPutError, c) {
		return
	}
	res := &v1.PutResponse{}
	res.Ok = true

	data, err = v1.EncodePutResponse(res)
	if CheckErrWithCode(err, EncodeError, c) {
		return
	}

	c.JSON(http.StatusOK, ApiResponse{
		Data: data,
		Code: 0,
	})
}
