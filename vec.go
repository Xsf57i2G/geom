package geom

import "math"

type Vec struct {
	X, Y, Z float64
}

func (a Vec) Add(b Vec) Vec {
	return Vec{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vec) Sub(b Vec) Vec {
	return Vec{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vec) Mul(b Vec) Vec {
	return Vec{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (a Vec) Div(b Vec) Vec {
	return Vec{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

func (a Vec) Scale(s float64) Vec {
	return Vec{a.X * s, a.Y * s, a.Z * s}
}

func (a Vec) Dot(b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec) Cross(b Vec) Vec {
	return Vec{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func (a Vec) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func (a Vec) Norm() Vec {
	return a.Scale(1 / a.Length())
}

func (a Vec) Project(b Vec) Vec {
	return b.Scale(a.Dot(b) / b.Dot(b))
}

func (a Vec) Reject(b Vec) Vec {
	return a.Sub(a.Project(b))
}

func (a Vec) Reflect(n Vec) Vec {
	return a.Reject(n.Scale(2 * a.Dot(n))).Norm()
}

func (a Vec) Lerp(b Vec, t float64) Vec {
	return a.Add(b.Sub(a).Scale(t))
}

func (a Vec) Angle(b Vec) float64 {
	return math.Acos(a.Dot(b) / (a.Length() * b.Length()))
}

func (a Vec) Dir(b Vec) Vec {
	return a.Sub(b).Norm()
}

func (a Vec) Dist(b Vec) float64 {
	return a.Sub(b).Length()
}
