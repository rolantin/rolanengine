package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
//To read all file content(in bytes) to memory, ioutil.ReadAll
func main()  {
	file,err := os.Open("RolanEngine/shader/Unlit.shader")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//// 使用ioutil读取文件所有内容
	b, err := ioutil.ReadAll(file)
	//fmt.Print(string(b))
	tracer := string(b)

	comma := strings.Index(tracer, "#ifdef FRAGMENT")
	//pos := strings.Index(string(b)[comma:], "#ifdef FRAGMENT")
	fmt.Println(tracer[comma:])
	//tracer[comma:]这个是的意思截取字符串tracer，从12开始，包括12
}

/*
//tracer := "死神来了,死神bye bye"
comma := strings.Index(tracer, ",")
//comma的意思是从字符串tracer查找第一个逗号，然后返回他的位置，这里的每个中文是占3个字符，从0开始计算，那么逗号的位置就是12

pos := strings.Index(tracer[comma:], "死神")
//tracer[comma:]这个是的意思截取字符串tracer，从12开始，包括12

fmt.Println(tracer[comma:])
//,死神bye bye
//整段的代码的意思是从tracer[comma:]这个字符串中查找“死神”这个字符串，第0位是逗号，第一位开始就是“死神”了，所以这里pos是1

fmt.Println(comma, pos, tracer[comma+pos+3:])*/