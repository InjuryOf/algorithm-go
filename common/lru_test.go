package common

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	lruCache := Construct(5)
	fmt.Printf("初始信息：%+v\n", lruCache)
	lruCache.Put(1, 10)
	fmt.Println("=====插入操作1=====")
	lruCache.getNodeInfo()
	lruCache.Put(2, 9)
	fmt.Println("=====插入操作2=====")
	lruCache.getNodeInfo()
	lruCache.Put(3, 8)
	fmt.Println("=====插入操作3=====")
	lruCache.getNodeInfo()
	lruCache.Put(4, 7)
	fmt.Println("=====插入操作4=====")
	lruCache.getNodeInfo()
	lruCache.Put(5, 6)
	fmt.Println("=====插入操作5=====")
	lruCache.getNodeInfo()
	// 该操作会触发内存淘汰
	lruCache.Put(6, 5)
	fmt.Println("=====插入操作6=====")
	lruCache.getNodeInfo()
	lruCache.Get(3)
	fmt.Println("=====查询操作=====")
	lruCache.getNodeInfo()
}
