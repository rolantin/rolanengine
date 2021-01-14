package math32

// Matrix3是一个内部组织为列矩阵的3x3矩阵
// Matrix3 is 3x3 matrix organized internally as column matrix
type Matrix3 [9]float32

// NewMatrix3  创建一个3x3 单位矩阵 creates and returns a pointer to a new Matrix3
// initialized as the identity matrix.
func NewMatrix3() *Matrix3 {
	var m Matrix3
	//创建一个单位矩阵
	m.Identity()
	return &m
}

// Set 设置3x3矩阵  set 3x3 column and Returns the pointer to this updated Matrix.
func (m *Matrix3) Set(x1, x2, x3, y1, y2, y3, z1, z2, z3 float32) *Matrix3 {
	m[0], m[1], m[2] = x1, y1, z1
	m[3], m[4], m[5] = x2, y2, z2
	m[6], m[7], m[8] = x3, y3, z3
	/*    | x1 y1 z1 |
	      | x2 y2 z2 |
	      | x3 y3 z3 |      */
	return m
}

// Identity 单位矩阵 sets this matrix as the identity matrix.
// Returns the pointer to this updated matrix.
func (m *Matrix3) Identity() *Matrix3 {
	m.Set(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)
	return m
}

// Zero 零矩阵 sets this matrix as the zero matrix.
// Returns the pointer to this updated matrix.
func (m *Matrix3) Zero() *Matrix3 {
	m.Set(
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	)
	return m
}

// Copy 复制源矩阵到源定义矩阵 copies src matrix into this one.
// Returns the pointer to this updated matrix.
func (m *Matrix3) Copy(src *Matrix3) *Matrix3 {
	*m = *src
	return m
}

//rolan to do
// MakeRotationFromQuaternion sets this matrix as a rotation matrix from the specified quaternion.
// Returns pointer to this updated matrix.
//https://blog.csdn.net/haima1998/article/details/80712257
func (m *Matrix3) MakeRotationFromQuaternion(q *Quaternion) *Matrix3{
	x := q.X
	y := q.Y
	z := q.Z
	w := q.W
	x2 := x + x
	y2 := y + y
	z2 := z + z
	xx := x * x2
	xy := x * y2
	xz := x * z2
	yy := y * y2
	yz := y * z2
	zz := z * z2
	wx := w * x2
	wy := w * y2
	wz := w * z2

	m[0] = 1 - (yy + zz)
	m[3] = xy - wz
	m[6] = xz + wy

	m[1] = xy + wz
	m[4] = 1 - (xx + zz)
	m[7] = yz - wx

	m[2] = xz - wy
	m[5] = yz + wx
	m[8] = 1 - (xx + yy)

	return m
}

//// ApplyToVector3Array multiplies length vectors in the array starting at offset by this matrix.
//// Returns pointer to the updated array.
//// This matrix is unchanged.
//func (m *Matrix3) ApplyToVector3Array(array []float32, offset int, length int) []float32 {
//	var v1 Vector3
//	j:= offset
//	for i:=0;i<length;i+=3 {
//		v1.X = array[j]
//		v1.Y = array[j+1]
//		v1.Z = array[j+2]
//		v1.
//	}
//	return array
//
//}