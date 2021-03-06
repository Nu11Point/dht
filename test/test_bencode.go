package main

import (
	"fmt"
	"github.com/owenliang/dht"
)

func main() {
	strData := "你好吗"
	if encData, err := dht.Encode(strData); err == nil {
		fmt.Println("encData=", string(encData))
	}

	intData := 1024
	if encData, err := dht.Encode(intData); err == nil {
		fmt.Println("encData=", string(encData))
	}

	listData := []interface{}{"你好吗", 1024,}
	if encData, err := dht.Encode(listData); err == nil {
		fmt.Println("encData=", string(encData))
	}

	dictData := map[string]interface{}{"t":"aa", "y":"q", "q":"ping", "a": map[string]interface{}{"id":"abcdefghij0123456789"}}
	if encData, err := dht.Encode(dictData); err == nil {
		fmt.Println("encData=", string(encData))
	}

	encIntData := []byte("i-12345e")
	if decData, err := dht.Decode(encIntData); err == nil {
		fmt.Println("decData=", decData)
	}

	encStrData := []byte("2:ab")
	if decData, err := dht.Decode(encStrData); err == nil {
		fmt.Println("decData=", decData)
	}

	encListData := []byte("l2:abl3:mmm1:ai5123eee")
	if decData, err := dht.Decode(encListData); err == nil {
		fmt.Println("decData=", decData)
	}

	encDictData := []byte("d2:abd2:cdl2:fgi5ed9:小电影i0eeeee")
	if decData, err := dht.Decode(encDictData); err == nil {
		fmt.Println("decData=", decData)
	}
}
