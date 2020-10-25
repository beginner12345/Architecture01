package main

// 缓存接口
type ICache interface {
	Get(k string) (v interface{})
	Set(k string, v interface{})
	Remove(k string)
	GetAmount() int
}
