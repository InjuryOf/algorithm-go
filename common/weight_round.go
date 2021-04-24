package common

type serverNode struct {
	name          string
	weight        int64
	currentWeight int64
}

func (sn *serverNode) incrWeight() {
	sn.currentWeight = sn.weight + sn.currentWeight
}

func (sn *serverNode) decrWeight(totalWeight int64) {
	sn.currentWeight = sn.currentWeight - totalWeight
}

func (sn *serverNode) setCurrentWeight(currentWeight int64) {
	sn.currentWeight = currentWeight
}

func weightedRound(serverList []*serverNode) *serverNode {
	var selectNode *serverNode
	var totalWeight, maxWeight int64
	for _, node := range serverList {
		node.incrWeight()
		if node.currentWeight > maxWeight{
			maxWeight = node.currentWeight
			selectNode = node
		}
		totalWeight += node.weight
	}

	if selectNode != (&serverNode{}) {
		selectNode.decrWeight(totalWeight)
	}

	return selectNode
}
