// Copyright 2020 just-coding-0.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"net/url"
	"strings"
)

func EncodePutRequest(request *PutRequest) (string, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return "", err
	}

	ret := url.QueryEscape(base64.URLEncoding.EncodeToString(data))
	ret = strings.Replace(ret, "%", "___", -1)
	return ret, nil
}

func DecodePutRequest(data string) (request *PutRequest, err error) {
	data = strings.Replace(data, "___", "%", -1)
	data, err = url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}

	_data, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	var r PutRequest
	if err := proto.Unmarshal(_data, &r); err != nil {
		return nil, err
	}
	request = &r
	return
}

func EncodePutResponse(request *PutResponse) (string, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return "", err
	}

	ret := url.QueryEscape(base64.URLEncoding.EncodeToString(data))
	ret = strings.Replace(ret, "%", "___", -1)
	return ret, nil
}

func DecodePutResponse(data string) (request *PutResponse, err error) {
	data = strings.Replace(data, "___", "%", -1)
	data, err = url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}

	_data, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	var r PutResponse
	if err := proto.Unmarshal(_data, &r); err != nil {
		return nil, err
	}
	request = &r
	return
}

func EncodeLoadRequest(request *LoadRequest) (string, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return "", err
	}

	ret := url.QueryEscape(base64.URLEncoding.EncodeToString(data))
	ret = strings.Replace(ret, "%", "___", -1)
	return ret, nil
}

func DecodeLoadRequest(data string) (request *LoadRequest, err error) {
	data = strings.Replace(data, "___", "%", -1)
	data, err = url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}

	_data, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	var r LoadRequest
	if err := proto.Unmarshal(_data, &r); err != nil {
		return nil, err
	}
	request = &r
	return
}

func EncodeLoadResponse(request *LoadResponse) (string, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return "", err
	}

	ret := url.QueryEscape(base64.URLEncoding.EncodeToString(data))
	ret = strings.Replace(ret, "%", "___", -1)
	return ret, nil
}

func DecodeLoadResponse(data string) (request *LoadResponse, err error) {
	data = strings.Replace(data, "___", "%", -1)
	data, err = url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}

	_data, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	var r LoadResponse
	if err := proto.Unmarshal(_data, &r); err != nil {
		return nil, err
	}
	request = &r
	return
}

func EncodeRemoveRequest(request *RemoveRequest) (string, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return "", err
	}

	ret := url.QueryEscape(base64.URLEncoding.EncodeToString(data))
	ret = strings.Replace(ret, "%", "___", -1)
	return ret, nil
}

func DecodeRemoveRequest(data string) (request *RemoveRequest, err error) {
	data = strings.Replace(data, "___", "%", -1)
	data, err = url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}

	_data, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	var r RemoveRequest
	if err := proto.Unmarshal(_data, &r); err != nil {
		return nil, err
	}
	request = &r
	return
}

func EncodeRemoveResponse(request *RemoveResponse) (string, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return "", err
	}

	ret := url.QueryEscape(base64.URLEncoding.EncodeToString(data))
	ret = strings.Replace(ret, "%", "___", -1)
	return ret, nil
}

func DecodeRemoveResponse(data string) (request *RemoveResponse, err error) {
	data = strings.Replace(data, "___", "%", -1)
	data, err = url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}

	_data, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	var r RemoveResponse
	if err := proto.Unmarshal(_data, &r); err != nil {
		return nil, err
	}
	request = &r
	return
}

func EncodePingRequest(request *PingRequest) (string, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return "", err
	}

	ret := url.QueryEscape(base64.URLEncoding.EncodeToString(data))
	ret = strings.Replace(ret, "%", "___", -1)
	return ret, nil
}

func DecodePingRequest(data string) (request *PingRequest, err error) {
	data = strings.Replace(data, "___", "%", -1)
	data, err = url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}

	_data, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	var r PingRequest
	if err := proto.Unmarshal(_data, &r); err != nil {
		return nil, err
	}
	request = &r
	return
}
