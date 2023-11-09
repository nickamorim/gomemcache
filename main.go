package main

import (
	"bytes"
	"fmt"

	"github.com/Shopify/gomemcache/memcache"
)

func main() {
	mc := memcache.New("udp://127.0.0.1:11211")
	v := bytes.Repeat([]byte("a"), 1024*1)
	// set keys foo-0 through foo-1000
	for i := 0; i < 1000; i++ {
		mc.Set(&memcache.Item{Key: fmt.Sprintf("foo-%d", i), Value: v})
	}
	// get key foo-0 through foo-1000
	for i := 0; i < 1000; i++ {
		res, err := mc.Get(fmt.Sprintf("foo-%d", i))
		if err != nil {
			fmt.Println("get")
			panic(err)
		}
		if string(v) != string(res.Value) {
			fmt.Println("v: ", v)
			fmt.Println("res.Value: ", res.Value)
			panic("value not equal")
		}
	}
	println("success...")
}
