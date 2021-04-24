package common

// 算法思路
// 1、当前权重计算：遍历对象的初始权重 + 当前权重
// 2、遍历对象的当前权重与最大权重值比较，如果大于最大权重值，则将当前对象置为选中对象，并将最大权重值设置为当前权重值。循环遍历直至循环结束
// 3、循环内总权重计算：遍历对象的初始权重值累加
// 4、循环结束后，如果判断是否存在选中对象，如果存在，则该对象的当前权重值 - 总权重

// 算法推导过程，模拟请求10次
// 请求1 {a-20,b-70,c-10} {0,0,0} -> b -> {20,-30,10}
// 请求2 {a-20,b-70,c-10} {40,40,20} -> a -> {-60,40,20}
// 请求3 {a-20,b-70,c-10} {-40,110,30} -> b-> {-40,10,30}
// 请求4 {a-20,b-70,c-10} {-20,80,40} -> b-> {-20,-20,40}
// 请求5 {a-20,b-70,c-10} {0,50,50} -> b-> {0,-50,50}
// 请求6 {a-20,b-70,c-10} {20,20,60} -> c-> {20,20,-40}
// 请求7 {a-20,b-70,c-10} {40,90,-30} -> b-> {40,-10,-30}
// 请求8 {a-20,b-70,c-10} {60,60,-20} -> a-> {-40,60,-20}
// 请求9 {a-20,b-70,c-10} {-20,130,-10} -> b-> {-20,30,-10}
// 请求10 {a-20,b-70,c-10} {0,100,0} -> b-> {0,0,0}



type serverNode struct {
	name          string
	weight        int64
	currentWeight int64
}

func (sn *serverNode) incrCurrentWeight() {
	sn.currentWeight = sn.currentWeight + sn.weight
}

func (sn *serverNode) decrCurrentWeight(totalWeight int64) {
	sn.currentWeight = sn.currentWeight - totalWeight
}

func weightRound(serverList []*serverNode) *serverNode {
	var maxWeight, totalWeight int64
	var selectNode *serverNode
	for _, node := range serverList {
		node.incrCurrentWeight()
		if node.currentWeight > maxWeight{
			selectNode = node
			maxWeight = node.currentWeight
		}
		totalWeight += node.weight
	}
	if selectNode != (&serverNode{}){
		selectNode.decrCurrentWeight(totalWeight)
	}
	return selectNode
}