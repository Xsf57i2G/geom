package geom

type Mat struct {
	Data []float64
}

func (m Mat) Mul(v Vec) Vec {
	return Vec{
		X: m.Data[0]*v.X + m.Data[1]*v.Y + m.Data[2]*v.Z,
	}
}
