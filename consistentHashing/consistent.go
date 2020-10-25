package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

//一致性哈希

type ConsistentHash struct {
	nodes map[string] ICache
	vnodes map[uint32] ICache
	vPerNode int
	circle []uint32
}

func NewConsistentHash(vPerNode int) *ConsistentHash {
	return &ConsistentHash {
		make(map[string] ICache),
		make(map[uint32] ICache),
		vPerNode,
		[]uint32{},
	}
}

func (ch *ConsistentHash) AddNode(name string, node ICache){
	if _, exist := ch.nodes[name]; exist{
		return
	}
	ch.nodes[name] = node
	for i:=1; i<=ch.vPerNode; i++ {
		vhash := vHashCode(name, i)
		ch.vnodes[vhash] = node
	}
	ch.sortCircle()
}

func (ch *ConsistentHash) RemoveNode(name string){
	delete(ch.nodes, name)
	for i:=1; i<=ch.vPerNode; i++ {
		vhash := vHashCode(name, i)
		delete(ch.vnodes, vhash)
	}
	ch.sortCircle()
}

func (ch *ConsistentHash) GetKV(k string) (v interface{}){
	node := ch.searchNode(k)
	return node.Get(k)
}

func (ch *ConsistentHash) SetKV(k string, v interface{}){
	node := ch.searchNode(k)
	node.Set(k, v)
}

func (ch *ConsistentHash) RemoveKV(k string){
	node := ch.searchNode(k)
	node.Remove(k)
}

func (ch *ConsistentHash) searchNode(k string) (node ICache){
	hash := hashCode(k)
	i := sort.Search(len(ch.circle), func(i int) bool { return ch.circle[i] >= hash })
	if i < len(ch.circle) {
		if i == len(ch.circle)-1 {
			i = 0
		}
	} else {
		i = len(ch.circle) - 1
	}
	return ch.vnodes[ch.circle[i]]
}

func (ch *ConsistentHash) sortCircle(){
	ch.circle = []uint32{}
	for k := range ch.vnodes {
		ch.circle = append(ch.circle, k)
	}
	sort.Slice(ch.circle, func(i, j int) bool {
		return ch.circle[i] < ch.circle[j]
	})
}

func hashCode(k string) uint32 {
	return crc32.ChecksumIEEE([]byte(k))
}

func vHashCode(name string, i int) uint32{
	vkey := fmt.Sprintln(name, "-", i)
	return hashCode(vkey)
}