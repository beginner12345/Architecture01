package main

func getCommonBetween2Lists(l1, l2 *ListNode)(*ListNode, bool){
	if l1==nil || l2==nil {return nil, false}
	len1, len2, p1, p2 := 0, 0, l1, l2
	//计算长度
	for p:=l1;p!=nil;p=p.Next {len1++}
	for p:=l2;p!=nil;p=p.Next {len2++}
	//两个链表先走到对齐长度
	if len1>=len2 {
		for i:=0;i<len1-len2;i++ {p1 = p1.Next}
	} else {
		for i:=0;i<len2-len1;i++ {p2 = p2.Next}
	}
	//比较公共节点
	for p1!= nil && p2!= nil {
		if p1 == p2 {return p1, true}
		p1, p2 = p1.Next, p2.Next
	}
	return nil, false
}