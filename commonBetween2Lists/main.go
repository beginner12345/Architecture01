package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	printResult(getCommonBetween2Lists(buildCase1()))
	printResult(getCommonBetween2Lists(buildCase2()))
	printResult(getCommonBetween2Lists(buildCase3()))
}


func printResult(first *ListNode, exist bool){
	if exist {
		fmt.Println("common node exist, first val =", first.Val)
	} else {
		fmt.Println("common node not exist")
	}
}

//构建测试链表
func buildListFrom(len int, from *ListNode) *ListNode{
	ret := from
	for i:=0;i<len;i++ {ret = &ListNode{i,ret}}
	return ret
}

func buildCase1() (l1, l2 *ListNode) {
	p := buildListFrom(4,nil)
	p1 := buildListFrom(3, p)
	p2 := buildListFrom(5, p)
	return p1, p2
}

func buildCase2() (l1, l2 *ListNode){
	p1 := buildListFrom(3, nil)
	p2 := buildListFrom(5, nil)
	return p1, p2
}

func buildCase3() (l1, l2 *ListNode){
	p1 := buildListFrom(5, nil)
	return p1, nil
}