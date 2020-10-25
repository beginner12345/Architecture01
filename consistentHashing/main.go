package main

import (
	"fmt"
	"math"
)

func main(){
	nodesCount := 10		//缓存节点数
	vCountPerNode := 100	//每个缓存节点的虚拟节点数
	amount := 1000000		//KV数量

	//创建缓存集群
	ch := NewConsistentHash(vCountPerNode)
	nodes := make([]ICache, nodesCount)
	for i:=0;i<nodesCount;i++{
		nodes[i] = NewListCache()
		ch.AddNode(fmt.Sprintln("node",i+1), nodes[i])
	}

	//向集群中添加数据
	SetKV(ch, amount)

	//计算标准差
	ComputeDeviation(nodes)
}

func SetKV(ch *ConsistentHash, amount int){
	for i:=0; i<amount;i++{
		ch.SetKV(fmt.Sprintln(i),"0") //主要统计KV分布性，v不重要
	}
}

func ComputeDeviation(nodes []ICache) {
	//打印每个节点的KV
	for i:=0;i<len(nodes);i++{
		fmt.Println("node",i+1,":",nodes[i].GetAmount())
	}
	//计算平均KV
	var ever float64 = 0.0
	for i:=range nodes {
		ever += float64(nodes[i].GetAmount())/float64(len(nodes))
	}
	fmt.Println("Everage: ", ever)

	//计算标准差
	var pever float64 = 0.0
	for i:=range nodes {
		pever += math.Pow(float64(nodes[i].GetAmount())-ever, 2)/float64(len(nodes))
	}
	dev := math.Sqrt(pever)
	fmt.Printf("Deviation: %2f", dev)
}