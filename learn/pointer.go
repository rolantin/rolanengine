package main
//http://tour.studygolang.com/moretypes/1
import "fmt"

type v3 struct {
	x float32
	y float32
}


func main() {
	//i, j := 42, 2701
	//
	//p := &i         // 指向 i
	//fmt.Println(*p) // 通过指针读取 i 的值
	//*p = 21         // 通过指针设置 i 的值
	//fmt.Println(i)  // 查看 i 的值
	//
	//p = &j         // 指向 j
	//*p = *p / 37   // 通过指针对 j 进行除法运算
	//fmt.Println(j) // 查看 j 的值
	var v v3
	p := &v.x //
	fmt.Println(p) //0xc0000a2058
	p = &v.y //
	fmt.Println(p) //0xc00000a0b4
	*p = 21
	fmt.Println(v.x,v.y) //0 21
	v.x = 10
	fmt.Println(v.x) //10

	vv := &v
	var ak = new(v3)
	ak.x = 15
	*vv = *ak
	println(&vv)
	print(vv.x)


}

type testvalue [2]float32

func (t *testvalue) wocao(a *testvalue) *testvalue{
    t = a
    return t
}
