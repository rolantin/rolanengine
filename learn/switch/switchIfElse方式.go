package main

import "fmt"

//switch 后面也可以不带表达式 类似if-else分支来使用

func main()  {
	var age int =10
	switch {
	case age==10:
		fmt.Println("age==10")
	case age==20:
		fmt.Println("age==20")
	default:
		fmt.Println("没有匹配到")
	}
//cast 中也可以对age的范围进行判断

	var score int = 30
	switch {
	case score<=10:
		fmt.Println("low")
	case score<=30 && score>=10:
		fmt.Println("med")
	default:
		fmt.Println("不及格")
	}

//switch 也可以定义一个变量  注意 定义变量得有分号！！！！！！！！！！！！！ 不建议这种写法
	switch ba:=50; {
	case ba>4:
		fmt.Println(">4")
	}
}