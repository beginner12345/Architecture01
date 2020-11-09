package main

import (
	"sort"
	"time"
)

type sortable []time.Duration
func (a sortable)Len() int {return len(a)}
func (a sortable)Swap(i, j int) {a[i],a[j] = a[j],a[i]}
func (a sortable)Less(i, j int) bool {return a[i]< a[j]}

func calculateResult(delays sortable) (ave, min, max, seq90, seq95, seq99 time.Duration) {
	sort.Sort(delays)
	//for i:=range delays{fmt.Println("delays", i, ":", delays[i])}
	min, max = delays[0], delays[len(delays)-1]
	i90, i95, i99 := len(delays)*90/100, len(delays)*95/100, len(delays)*99/100
	seq90, seq95, seq99 = delays[i90-1], delays[i95-1], delays[i99-1]
	ave = average(delays)
	return
}

func average(a []time.Duration) time.Duration{
	var ret time.Duration
	for i := range a {
		ret += a[i] / time.Duration(len(a))
	}
	return ret
}

