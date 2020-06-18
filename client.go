// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package go_cache

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/just-coding-0/go-cache/internal/bytesconv"
	"github.com/just-coding-0/go-cache/internal/consistent_hashing"
	"github.com/just-coding-0/go-cache/v1"
	"io/ioutil"
	"log"
	m_rand "math/rand"
	"net/http"
	"time"
)

const (
	ping      = "/ping"
	mapPut    = "/map/put"
	mapRemove = "/map/remove"
	mapLoad   = "/map/load"
)

var (
	putValueFail = errors.New("put value fail")
)

type client struct {
	consistent *consistent_hashing.Consistent
	PublicKey  *rsa.PublicKey
	Nodes      []*consistent_hashing.Node
	exitChan   chan struct{}
}

type ApiResponse struct {
	Msg  string `json:"err,omitempty"`
	Data string `json:"data,omitempty"`
	Code int    `json:"code,omitempty"`
}

func NewClient(nodes []*consistent_hashing.Node, PublicKeyPath string) (Client, error) {
	c := &client{}
	c.exitChan = make(chan struct{}, 1)
	c.consistent = consistent_hashing.NewConsistent()
	c.Nodes = nodes
	for _, node := range nodes {
		c.consistent.Add(node)
	}
	pubKey, err := GetPublicKey(PublicKeyPath)
	if err != nil {
		return nil, err
	}
	c.PublicKey = pubKey
	c.ping()
	go c.PingUpstreamService()

	return c, nil
}

type Client interface {
	MapPut(key, value string, expiryPolicy int64) error // set key value
	MapLoad(key string) (string, bool, error)           // get value
	MapRemove(key string) (bool, error)                 // remove

	PingUpstreamService() // ping 上游服务器

	Close() // 释放资源
}

func (c *client) MapPut(key, value string, expiryPolicy int64) error {
	node := c.consistent.Get(key)
	encodeStr, _ := v1.EncodePutRequest(&v1.PutRequest{Key: key, Value: value, ExpiryPolicy: expiryPolicy})
	requestUrl := fmt.Sprintf("%s:%d%s?request=%s", node.Ip, node.Port, mapPut, encodeStr)
	res, err := http.Get(requestUrl)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	var apiRes ApiResponse
	if err = json.Unmarshal(buf, &apiRes); err != nil {
		return err
	}
	var putRes *v1.PutResponse
	putRes, err = v1.DecodePutResponse(apiRes.Data)
	if err != nil {
		return err
	}

	if putRes.Ok == false {
		return putValueFail
	}

	return nil
}

func (c *client) MapLoad(key string) (string, bool, error) {
	node := c.consistent.Get(key)
	encodeStr, _ := v1.EncodeLoadRequest(&v1.LoadRequest{Key: key})

	requestUrl := fmt.Sprintf("%s:%d%s?request=%s", node.Ip, node.Port, mapLoad, encodeStr)
	res, err := http.Get(requestUrl)
	if err != nil {
		return "", false, nil
	}

	buf, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", false, err
	}

	var apiRes ApiResponse
	if err = json.Unmarshal(buf, &apiRes); err != nil {
		return "", false, err
	}
	var loadRes *v1.LoadResponse
	loadRes, err = v1.DecodeLoadResponse(apiRes.Data)
	if err != nil {
		return "", false, err
	}

	return loadRes.Value, loadRes.Ok, nil
}

func (c *client) MapRemove(key string) (bool, error) {
	node := c.consistent.Get(key)

	encodeStr, _ := v1.EncodeRemoveRequest(&v1.RemoveRequest{Key: key})

	requestUrl := fmt.Sprintf("%s:%d%s?request=%s", node.Ip, node.Port, mapRemove, encodeStr)
	res, err := http.Get(requestUrl)
	if err != nil {
		return false, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return false, err
	}
	fmt.Printf("%s \n", buf)

	return true, nil
}

func (c *client) Close() {
	c.exitChan <- struct{}{}
}

func (c *client) PingUpstreamService() {
	t := time.NewTicker(time.Hour * 6)
	for {
		select {
		case <-t.C:
			c.ping()
		case <-c.exitChan:
			t.Stop()
			close(c.exitChan)
			return
		}
	}
}

func (c *client) ping() {
	for idx := range c.Nodes {
		if c.Nodes[idx] == nil {
			continue
		}

		plaintext := fmt.Sprintf("%d", m_rand.Uint64())
		cipher, err := rsa.EncryptPKCS1v15(rand.Reader, c.PublicKey, bytesconv.StringToBytes(plaintext))

		if err != nil {
			panic(err)
		}

		pingRequest := &v1.PingRequest{}
		encodeCipher := base64.URLEncoding.EncodeToString(cipher)

		pingRequest.Cipher = encodeCipher
		pingRequest.Plaintext = plaintext

		encodeStr, err := v1.EncodePingRequest(pingRequest)
		if err != nil {
			log.Println(err)
		}
		requestUrl := fmt.Sprintf("%s:%d%s?request=%s", c.Nodes[idx].Ip, c.Nodes[idx].Port, ping, encodeStr)

		res, err := http.Get(requestUrl)
		if err != nil {
			log.Println(err)
		}
		buf, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		var api ApiResponse
		json.Unmarshal(buf, &api)

		if api.Code != 0 {
			c.consistent.Remove(c.Nodes[idx])
			c.Nodes[idx] = nil
		}
	}
}
