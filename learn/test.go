package main

type num struct {
	x float32
	y float32

}

type v3 struct {
	x float32
	y float32
	z float32
}

func (n *num) ass(value float32)  {
	n.x = value
}

func (v *v3) add(other *v3) *v3  {
	*v = *other

	//v.x += other.x
	//v.y += other.y
	//v.z += other.z
	return v
}

func (v *v3) copy(other *v3) *v3  {
	v.x  = other.x
	v.y  = other.y
	v.z  = other.z

	*v = *other
	return v
}

func main()  {
	var v v3
	v.add(&v3{1,2,3})

	//v.copy(&v3{2,2,2})
	//println(v.x,v.y,v.z)


	//var cc num
	//cc.ass(3)
   //println(cc.x)
}



