package common

import "fmt"

// 双向链表
type LinkNode struct {
	key, value int
	pre, next  *LinkNode
}

//
type LRUCache struct {
	hm         map[int]*LinkNode //指向链表的指针
	cap        int               //长度
	head, tail *LinkNode
}

// 构建LRU结构
func Construct(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

// 查询
func (lr *LRUCache) Get(key int) int {
	cache := lr.hm
	if node, exists := cache[key]; exists {
		lr.moveToHead(node)
		return node.value
	}
	return -1
}

// 插入
func (lr *LRUCache) Put(key, value int) int {
	head := lr.head
	tail := lr.tail
	cache := lr.hm
	if node, exists := cache[key]; exists {
		node.value = value
		lr.moveToHead(node)
		return 1
	} else {
		node := &LinkNode{key, value, nil, nil}
		//判断缓存容量是否已满
		if len(cache) >= lr.cap {
			delete(cache, tail.pre.key)
			tail.pre.pre.next = tail
			tail.pre = tail.pre.pre
		}
		node.next = head.next
		node.pre = head
		head.next.pre = node
		head.next = node
		cache[key] = node
		return 1
	}
	return 0
}

// 删除key
func (lr *LRUCache)Del(key int) int  {
	cache := lr.hm
	if node, exists := cache[key]; exists {
		delete(cache, key)
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	return 0
}

func (lr *LRUCache) moveToHead(node *LinkNode) int {
	head := lr.head
	//从当前位置删除节点
	node.pre.next = node.next
	node.next.pre = node.pre
	//将节点插入到头部
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
	return -1
}

func (lr *LRUCache) getNodeInfo() {
	for _, node := range lr.hm {
		fmt.Printf("缓存头尾信息：%v=%v  ||   节点信息：node-value-%d >> node-pre-%d=next-%d\n", lr.head.next.key, lr.tail.pre.key, node.key, node.pre.key, node.next.key)
	}
}
