// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package handler

import (
	"crypto/rsa"
	"github.com/gin-gonic/gin"
	"github.com/just-coding-0/go-cache"
	"net/http"
	"sync"
)

var (
	Manager    go_cache.CacheManager
	IpMap      *sync.Map
	PrivateKey *rsa.PrivateKey
)

type StatusCode int

const (
	SUCCESS StatusCode = 0

	DecodeError StatusCode = 1001
	EncodeError StatusCode = 1002
	MapPutError StatusCode = 2001

	RsaDecryptError StatusCode = 4001
	RsaEncryptError StatusCode = 4002

	RasCipherTextUnValidError StatusCode = 4003
)

var (
	decodeErrorMsg = "解码失败"
	encodeErrorMsg = "编码失败"

	mapPutErrorMsg = "map put error"

	RsaDecryptErrorMsg           = "解密失败"
	rasCipherTextUnValidErrorMsg = "密文错误"
)

type ApiResponse struct {
	Msg  string     `json:"err,omitempty"`
	Data string     `json:"data,omitempty"`
	Code StatusCode `json:"code,omitempty"`
}

func getErrorByCode(code StatusCode) string {
	switch code {
	case DecodeError:
		return decodeErrorMsg
	case EncodeError:
		return encodeErrorMsg
	case MapPutError:
		return mapPutErrorMsg
	case RsaDecryptError:
		return RsaDecryptErrorMsg
	case RasCipherTextUnValidError:
		return rasCipherTextUnValidErrorMsg
	default:
		return ""
	}
}

func SetCacheManager(mem uint64, shardSize uintptr, mode uint8) func() {
	manager := go_cache.NewManager(mem, shardSize, mode)
	Manager = manager

	return func() {
		manager.Close()
	}
}

func SetPrivateKey(privateKeyPath string) error {
	key, err := go_cache.GetPrivateKey(privateKeyPath)
	if err != nil {
		return err
	}
	PrivateKey = key
	return nil
}

func SetIpMap(_ipMap *sync.Map) {
	IpMap = _ipMap
}
func CheckErrWithCode(err error, code StatusCode, c *gin.Context) (ref bool) {
	ref = err != nil
	if ref {
		c.JSON(http.StatusOK, ApiResponse{
			Code: code,
			Msg:  err.Error(),
			Data: "",
		})
	}
	return
}

func CheckWithCode(ok bool, code StatusCode, c *gin.Context) (ref bool) {
	ref = ok
	if ref {

		c.JSON(http.StatusOK, ApiResponse{
			Code: code,
			Data: "",
			Msg:  getErrorByCode(code),
		})
	}
	return
}
