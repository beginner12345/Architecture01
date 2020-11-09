package main

import (
	"fmt"
	"os"
	"strconv"
)


func main(){
	//获取参数
	url, con, num := analyzeArgs()
	if con<=0 || num <=0 {return}

	//创建并发任务
	result := scheduleTask(url, con, num)

	//统计结果
	ave, min, max, seq90, seq95, seq99 := calculateResult(result)
	fmt.Println("ave:", ave, "min:", min, "max:", max, "seq90:", seq90, "seq95:", seq95, "seq99:", seq99)
}

//命令行执行 参数解析
func analyzeArgs() (url string, con, num int){
	for i:=1; i<len(os.Args);i=i+2 {
		if os.Args[i] == "-u" {
			url = os.Args[i+1]
		} else if os.Args[i] == "-c" {
			con, _ = strconv.Atoi(os.Args[i+1])
		} else if os.Args[i] == "-n" {
			num, _ = strconv.Atoi(os.Args[i+1])
		}
	}
	return
}

//IDE执行 默认参数
func analyzeArgs1() (url string, con, num int){
	return "https://www.baidu.com", 10, 100
}