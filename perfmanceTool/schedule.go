package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func scheduleTask(url string, countConcur, countReq int) []time.Duration {
	//结果缓存
	ret := []time.Duration{}
	//并发监听
	var wg sync.WaitGroup
	wg.Add(countConcur)
	//数据通道
	ch := make(chan []time.Duration, countConcur)
	//将请求分摊给每个并发
	counts := make([]int, countConcur)
	for i := range counts {
		counts[i] = countReq/countConcur
	}
	for i:=0; i < countReq%countConcur; i++ {
		counts[i]++
	}
	//先执行一次,避免tcp建立影响延迟,不计入统计结果
	if _, err := http.Get(url); err != nil{
		fmt.Println("main http err: ", err)
		return ret
	}
	//启动并等待并发完成
	for i := range counts {
		go run(url, counts[i], &wg, ch)
	}
	wg.Wait()
	//合并结果
	for i:=0;i<countConcur;i++{
		t, ok := <-ch
		if ok {ret = append(ret, t...)}
	}
	return ret
}

func run(url string, count int, wg *sync.WaitGroup, ch chan[]time.Duration) {
	defer wg.Done()
	delays := make([]time.Duration, count)
	//访问并记录延迟,通过chan返回结果
	for i:=0; i<count; i++ {
		t0 := time.Now()
		if _, err := http.Get(url); err != nil{
			fmt.Println("sub http err: ", err)
			return
		}
		delays[i] = time.Since(t0)
	}
	ch <- delays
}
