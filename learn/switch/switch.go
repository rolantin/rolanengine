package main

import "fmt"

/*
1.case 后面是一个常量 变量 也可以有返回值的函数 或者运算
2.switch的数据类型 要和 case的数据类型一样
3.switch允许多个表达式
4.switch不允许相同的表达式 比如 cast 5  第二个又是 cast 5 但是使用变量可以，骗过编译器
5.default 语句可以不写
 */


func main()  {
	//golang 是不需要break
	//定义一个变量接受字符
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d")
	fmt.Scanf("%c",&key)
	//base(key)
	base(test(key)+1) //注意！这里回返回 c 0 +1 +1 等于第三个 所以为输出 c
	//多个表达式
	var n1 int
	switch n1 {
	case 10,20:
		fmt.Println("ok")
	default:
		fmt.Println("end")
	}
}

//使用函数作为switch的值
func test(b byte) byte{ //理解这种语法现象
	return b+1
}

func base(key byte)  {
	switch key {
	case 'a':
		fmt.Println("you input a")
	case 'b':
		fmt.Println("you input b")
	case 'c':
		fmt.Println("you input c")
	case 'd':
		fmt.Println("you input d")
	default:
		fmt.Println("you input wrong")
	}
}