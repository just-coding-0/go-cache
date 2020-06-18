package go_cache

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

func GetPrivateKey(priKeyPath string) (*rsa.PrivateKey, error) {
	file, err := os.Open(priKeyPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("invalid rsa private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func GetPublicKey(publicKeyPath string) (publicKey *rsa.PublicKey, err error) {

	file, err := os.Open(publicKeyPath)
	if err != nil {
		return
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("invalid rsa public key")
	}

	_publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	publicKey = _publicKey.(*rsa.PublicKey)
	return
}

func Json(obj interface{}) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.Encode(obj)

	return buf.Bytes()
}