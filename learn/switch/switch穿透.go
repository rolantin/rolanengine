package main

import "fmt"

func main()  {
	age:=10
	switch age {
	case 5: // --> 假如想这样做穿透，其实可以这样写 case 5,10
		fmt.Println("good")
		fallthrough // --> 这个可以穿透下一层 第三层就不行了
	case 10:
		fmt.Println("good2")
	}
}