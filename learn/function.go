package main

type Vector3 struct {
	X float32
	Y float32
	Z float32
}
//成员方法 得 m.new
func (m *Vector3)NewVec3() *Vector3{
	m.X =0
	m.Y =0
	m.Z =0
	return m
}

//全局方法
// NewVector3 creates and returns a pointer to a new Vector3 with
// the specified x, y and y components
func NewVector3(x,y,z float32) *Vector3 {
	return &Vector3{X:x,Y:y,Z:z}
}