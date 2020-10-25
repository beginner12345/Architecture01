package main

import (
	"container/list"
	)

// 自定义cache

//主要测试一致性哈希的算法均衡性，使用链表替代缓存，只保存key
type ListCache struct {
	cache *list.List
}

func NewListCache() *ListCache {
	return &ListCache{list.New()}
}

func (c *ListCache) Get(k string) (v interface{}){
	return v
}

func (c *ListCache) Set(k string, v interface{}) {
	c.cache.PushBack(k)
}

func (c *ListCache) Remove(k string) {

}

func (c *ListCache) GetAmount() int{
	return c.cache.Len()
}
