package baseMesh

type RoBaseMesh struct {
	triangle []float32
}

func(f *RoBaseMesh) CreateTriangle() []float32  {
     f.triangle = []float32{
		 0, 0.5, 0, // top
		 -0.5, -0.5, 0, // left
		 0.5, -0.5, 0, // right
	 }
	 return f.triangle
}

