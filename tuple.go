package raytracer

import (
	"fmt"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func TupleEqual(a, b Tuple) bool {
	x := FloatEqual(a.X, b.X)
	y := FloatEqual(a.Y, b.Y)
	z := FloatEqual(a.Z, b.Z)
	w := FloatEqual(a.W, b.W)
	return (x && y && z && w)
}

func (a Tuple) Add(b Tuple) (Tuple, error) {
	if a.IsPoint() && b.IsPoint() {
		return Tuple{}, fmt.Errorf("Can not add a point to a point: a = %+v, b = %+v", a, b)
	}

	return Tuple{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z, W: a.W + b.W}, nil
}

func (a Tuple) Subtract(b Tuple) (Tuple, error) {
	if a.IsVector() && b.IsPoint() {
		return Tuple{}, fmt.Errorf("Can not subtract a point from a vector: a = %+v, b = %+v", a, b)
	}

	return Tuple{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z, W: a.W - b.W}, nil
}

func (a Tuple) Negate() Tuple {
	return Tuple{X: a.X * -1, Y: a.Y * -1, Z: a.Z * -1, W: a.W * -1}
}

func FromMatrix(m Matrix) (Tuple, error) {
	if m.Rows != 4 || m.Cols != 1 {
		return Tuple{}, fmt.Errorf("Can not form a tuple from matrix: %+v", m)
	}

	x := m.Data[0][0]
	y := m.Data[1][0]
	z := m.Data[2][0]
	w := m.Data[3][0]
	return Tuple{X: x, Y: y, Z: z, W: w}, nil
}

func (a Tuple) MultiplyFloat(f float64) Tuple {
	return Tuple{X: a.X * f, Y: a.Y * f, Z: a.Z * f, W: a.W * f}
}

func (a Tuple) MultiplyTuple(b Tuple) Tuple {
	return Tuple{X: a.X * b.X, Y: a.Y * b.Y, Z: a.Z * b.Z, W: a.W * b.W}
}

func (a Tuple) Magnitude() float64 {
	sum := math.Pow(a.X, 2) + math.Pow(a.Y, 2) + math.Pow(a.Z, 2) + math.Pow(a.W, 2)
	return math.Sqrt(sum)
}

func (a Tuple) Normalize() (Tuple, error) {
	m := a.Magnitude()
	if m == 0 {
		return Tuple{}, fmt.Errorf("Can not divide by zero: %+v", a)
	}
	return Tuple{X: a.X / m, Y: a.Y / m, Z: a.Z / m, W: a.W / m}, nil
}

func Dot(t1, t2 Tuple) (float64, error) {
	if t1.IsPoint() || t2.IsPoint() {
		return 1.0, fmt.Errorf("Can not perform dot product on points: t1: %+v, t2: %+v", t1, t2)
	}

	return (t1.X * t2.X) + (t1.Y * t2.Y) + (t1.Z * t2.Z) + (t1.W * t2.W), nil
}

func Cross(t, b Tuple) (Tuple, error) {
	if t.IsPoint() || b.IsPoint() {
		return Tuple{}, fmt.Errorf("can not perform cross product on points: t: %+v, b: %+v", t, b)
	}

	zy := t.Y*b.Z - t.Z*b.Y
	zx := t.Z*b.X - t.X*b.Z
	xy := t.X*b.Y - t.Y*b.X
	return Tuple{X: zy, Y: zx, Z: xy, W: 0.0}, nil
}

func Reflect(in, norm Tuple) (Tuple, error) {
	if in.IsPoint() || norm.IsPoint() {
		return Tuple{}, fmt.Errorf("can not reflect points: in: %+v, norm: %+v", in, norm)
	}

	d, err := Dot(in, norm)
	if err != nil {
		return Tuple{}, err
	}

	r, err := in.Subtract(norm.MultiplyFloat(2 * d))
	if err != nil {
		return Tuple{}, err
	}

	return r, nil
}
