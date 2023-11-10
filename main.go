package main

import (
	"bytes"
	"fmt"

	"github.com/Shopify/gomemcache/memcache"
)

func main() {
	mc := memcache.New("udp://127.0.0.1:12312")
	v := bytes.Repeat([]byte("a"), 8*1024)
	for i := 0; i < 1; i++ {
		err := mc.Set(&memcache.Item{Key: fmt.Sprintf("bar-%d", i), Value: v})
		if err != nil {
			fmt.Println("set")
			panic(err)
		}
	}
	for i := 0; i < 1; i++ {
		res, err := mc.Get(fmt.Sprintf("bar-%d", i))
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
	fmt.Println("success...")
}
