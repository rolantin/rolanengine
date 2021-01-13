package config

const (
	WindowWidth = 1024
	WindowHeight = 768
)

//const (
//	windowWidth  = 1024
//	windowHeight = 768
//)


//Golang 中结构体常量的三种实现方式
/*
type User interface {
	Name() string
}

type a struct {}
func (a *a) Name() string {
	return "xiaoming"
}

type b struct {
	name string
}

func (b *b) Name() string {
	return b.name
}

type c struct {
	Name string
}

var ConstA = a{}
var ConstB = b{ name: "xiaoming" }
var ConstC = c{ Name: "xiaoming" }
*/
