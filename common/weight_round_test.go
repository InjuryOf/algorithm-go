package common

import (
	"fmt"
	"testing"
)

func TestWeightRound(t *testing.T) {
	// 初始化服务器信息
	var serverList []*serverNode
	serverList = append(serverList,
		&serverNode{
			name:          "服务器A",
			weight:        20,
			currentWeight: 0,
		}, &serverNode{
			name:          "服务器B",
			weight:        70,
			currentWeight: 0,
		}, &serverNode{
			name:          "服务器C",
			weight:        10,
			currentWeight: 0,
		})

	// 模拟请求
	for i := 1; i <= 10; i++ {
		currentNode := weightedRound(serverList)
		fmt.Printf("请求次数：%d 当前节点：%s-分配权重（%d）- 当前权重（%d）\n", i,
		currentNode.name, currentNode.weight, currentNode.currentWeight)
	}
}
