package common

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	lruCache := Construct(5)
	fmt.Printf("初始信息：%+v\n", lruCache)

	fmt.Println("=====插入操作1=====", lruCache.Put(1, 10))
	lruCache.getNodeInfo()
	fmt.Println("=====插入操作2=====", lruCache.Put(2, 9))
	lruCache.getNodeInfo()
	fmt.Println("=====插入操作3=====", lruCache.Put(3, 8))
	lruCache.getNodeInfo()
	fmt.Println("=====插入操作4=====", lruCache.Put(4, 7))
	lruCache.getNodeInfo()
	fmt.Println("=====插入操作5=====", lruCache.Put(5, 6))
	lruCache.getNodeInfo()
	// 该操作会触发内存淘汰
	fmt.Println("=====插入操作6=====", lruCache.Put(6, 5))
	lruCache.getNodeInfo()
	fmt.Println("=====查询操作=====", lruCache.Get(3))
	lruCache.getNodeInfo()
	fmt.Println("=====删除操作=====", lruCache.Del(4))
	lruCache.getNodeInfo()
}
