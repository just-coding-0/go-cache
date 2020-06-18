package handler

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/just-coding-0/go-cache/internal/bytesconv"
	v1 "github.com/just-coding-0/go-cache/v1"
	"net/http"
	"time"
)

func PingHandler(c *gin.Context) {

	data := c.Query(`request`)

	pingRequest, err := v1.DecodePingRequest(data)
	if CheckErrWithCode(err, DecodeError, c) {
		return
	}

	cipher, _ := base64.URLEncoding.DecodeString(pingRequest.Cipher)
	buf, err := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, cipher)
	if CheckErrWithCode(err, RsaDecryptError, c) {
		return
	}

	_plaintext := bytesconv.BytesToString(buf)

	if CheckWithCode(pingRequest.Plaintext != _plaintext, RasCipherTextUnValidError, c) {
		return
	}

	IpMap.Store(c.ClientIP(), time.Now().Format(`2006-01-02`))

	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
