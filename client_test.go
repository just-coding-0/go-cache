package go_cache

import (
	"github.com/go-playground/assert/v2"
	"github.com/just-coding-0/go-cache/internal/consistent_hashing"
	"log"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	node := &consistent_hashing.Node{}
	node.Id = 1
	node.Ip = "http://127.0.0.1"
	node.Port = 10001
	node.Weight = 1
	client, err := NewClient([]*consistent_hashing.Node{node}, "/Users/ted/go/src/github.com/just-coding-0/go-cache/rsa_public_key.pem")
	if err != nil {
		t.Fatal(err)
	}
	var keys = []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7", "key8"}
	var values = []string{"value1", "value2", "value3", "value4", "value5", "value6", "value7", "value8"}

	for i := 0; i < len(keys); i++ {
		err = client.MapPut(keys[i], values[i], time.Now().Unix()+60)
		if err != nil {
			log.Println(err)
		}
	}

	for i := 0; i < len(values); i++ {
		val, ok, err := client.MapLoad(keys[i])
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, ok, true)
		assert.Equal(t, val, values[i])
	}

}


