package math32

// Vector3 is a 3D vector/point with X, Y and Z components.
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

//func (m *Vector3)NewVec3() *Vector3{
//	m.X =0
//	m.Y =0
//	m.Z =0
//	return m
//}


// NewVector3 creates and returns a pointer to a new Vector3 with
// the specified x, y and y components
func NewVector3(x,y,z float32) *Vector3 {
	return &Vector3{X:x,Y:y,Z:z}
}

// NewVec3 creates and returns a pointer to a new zero-ed Vector3.
func NewVec3() *Vector3{
	return &Vector3{X:0,Y:0,Z:0}
}

// Set sets this vector X, Y and Z components.
// Returns the pointer to this updated vector.
func (v *Vector3) Set(x,y,z float32) *Vector3 {
	v.X,v.Y,v.Z = x,y,z
	return v
}

// SetX sets this vector X component.
// Returns the pointer to this updated Vector.
func (v *Vector3) SetX(x float32) *Vector3  {
	v.X = x
	return v
}

// SetY sets this vector Y component.
// Returns the pointer to this updated vector.
func (v *Vector3) SetY(y float32) *Vector3 {
	v.Y = y
	return v
}

// SetZ sets this vector Z component.
// Returns the pointer to this updated vector.
func (v *Vector3) SetZ(z float32) *Vector3 {
	v.Z = z
	return v
}

// SetComponent sets this vector component value by its index: 0 for X, 1 for Y, 2 for Z.
// Returns the pointer to this updated vector
func (v *Vector3) SetCompont(index int,value float32) {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	default:
		panic("index is out of range: max is 0-2")
	}
}

// Component 返回某个分量 returns this vector component by its index: 0 for X, 1 for Y, 2 for Z.
func (v *Vector3) Component(index int) float32 {
	switch index {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	default:
		panic("index is out of range: max is 0-2")
	}
}

// SetByName sets this vector component value by its case insensitive name: "x", "y", or "z".
func (v *Vector3) SetByName(name string,value float32){
	switch name {
	case "x","X":
		v.X = value
	case "y","Y":
		v.Y = value
	case "z","Z":
		v.Z = value
	default:
		panic("Invalid Vector3 component name: " + name)
	}
}

// Zero sets this vector X, Y and Z components to be zero.
// Returns the pointer to this updated vector.
func (v *Vector3) Zero() *Vector3  {
	v.X,v.Y,v.Z=0,0,0
	return v
}

// Copy copies other vector to this one.
// It is equivalent to: *v = *other.
// Returns the pointer to this updated vector.
func (v *Vector3) Copy(v3 *Vector3) *Vector3 {
	*v = *v3
	return v
}

// Add adds other vector to this one.
// Returns the pointer to this updated vector.
func (v *Vector3) Add(v3 *Vector3) *Vector3  {
	v.X += v3.X
	v.Y += v3.Y
	v.Y += v3.Z
	return v
}


