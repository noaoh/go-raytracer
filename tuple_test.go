package raytracer 

import (
	"math"
	"testing"
)

func TestIsPoint(t *testing.T) {
	expected := []bool{true, true, false, false}
	actual := []Tuple{
		Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
		Tuple{X: 1.0, Y: 1.0, Z: 1.0, W: 1.0},
		Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 0.0},
		Tuple{X: 1.0, Y: 1.0, Z: 1.0, W: 0.0},
	}

	for i, x := range actual {
		t.Run(string(i), func(t *testing.T) {
			if x.IsPoint() != expected[i] {
				t.Logf("Failed IsPoint for: %+v\n", x)
				t.Fail()
			}
		})
	}
}

func TestIsVector(t *testing.T) {
	expected := []bool{true, true, false, false}
	actual := []Tuple{
		Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
		Tuple{X: 1.0, Y: 1.0, Z: 1.0, W: 1.0},
		Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 0.0},
		Tuple{X: 1.0, Y: 1.0, Z: 1.0, W: 0.0},
	}

	for i, x := range actual {
		t.Run(string(i), func(t *testing.T) {
			if x.IsPoint() != expected[i] {
				t.Logf("Failed IsVector for: %+v\n", x)
				t.Fail()
			}
		})
	}
}

func TestTupleAdd(t *testing.T) {
	x := Tuple{X: 3.0, Y: -2.0, Z: 5.0, W: 1.0}
	y := Tuple{X: -2.0, Y: 3.0, Z: 1.0, W: 0.0}
	e := Tuple{X: 1.0, Y: 1.0, Z: 6.0, W: 1.0}
	a, _ := x.Add(y)
	if a != e {
		t.Logf("Add(%+v, %+v) != %+v\n", x, y, e)
		t.Fail()
	}
}

func TestTupleSubtract(t *testing.T) {
	expected := []Tuple{
		Tuple{X: -2.0, Y: -4.0, Z: -6.0, W: 0.0},
		Tuple{X: -2.0, Y: -4.0, Z: -6.0, W: 1.0},
		Tuple{X: -2.0, Y: -4.0, Z: -6.0, W: 0.0},
	}

	actual := [][]Tuple{
		{Tuple{X: 3.0, Y: 2.0, Z: 1.0, W: 1.0}, Tuple{X: 5.0, Y: 6.0, Z: 7.0, W: 1.0}},
		{Tuple{X: 3.0, Y: 2.0, Z: 1.0, W: 1.0}, Tuple{X: 5.0, Y: 6.0, Z: 7.0, W: 0.0}},
		{Tuple{X: 3.0, Y: 2.0, Z: 1.0, W: 0.0}, Tuple{X: 5.0, Y: 6.0, Z: 7.0, W: 0.0}},
	}

	for i, e := range expected {
		t.Run(string(i), func(t *testing.T) {
			x := actual[i][0]
			y := actual[i][1]
			s, err := x.Subtract(y); if err != nil {
                                t.Log(err)
                                t.Fail()
                        }
			if s != e {
				t.Logf("Subtract(%+v,%+v) != %+v\n", x, y, e)
				t.Fail()
			}
		})
	}
}

func TestNegate(t *testing.T) {
	e := Tuple{X: -1, Y: 2, Z: -3, W: 4}
	a := Tuple{X: 1, Y: -2, Z: 3, W: -4}
	if a.Negate() != e {
		t.Logf("Negate(%+v) != %+v", a, e)
	}
}

func TestMultiplyFloat(t *testing.T) {
	f := []float64{3.5, 0.5}
	a := Tuple{X: 1.0, Y: -2.0, Z: 3.0, W: -4.0}
	e := []Tuple{
		Tuple{X: 3.5, Y: -7.0, Z: 10.5, W: -14.0},
		Tuple{X: 0.5, Y: -1.0, Z: 1.5, W: -2.0},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			if a.MultiplyFloat(f[i]) != x {
				t.Logf("Multiply(%+v, %f) != %+v\n", a, f[i], x)
				t.Fail()
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	a := []Tuple{
		Tuple{X: 1.0, Y: 0.0, Z: 0.0, W: 0.0},
		Tuple{X: 0.0, Y: 1.0, Z: 0.0, W: 0.0},
		Tuple{X: 0.0, Y: 0.0, Z: 1.0, W: 0.0},
		Tuple{X: 1.0, Y: 2.0, Z: 3.0, W: 0.0},
		Tuple{X: -1.0, Y: -2.0, Z: -3.0, W: 0.0},
	}
	e := []float64{1, 1, 1, math.Sqrt(14), math.Sqrt(14)}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			if a[i].Magnitude() != x {
				t.Logf("Magnitude(%+v) != %f\n", a[i], x)
				t.Fail()
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	a := Tuple{X: 1, Y: 2, Z: 3, W: 0}
	e := Tuple{X: 1 / math.Sqrt(14), Y: 2 / math.Sqrt(14), Z: 3 / math.Sqrt(14), W: 0}
	n, _ := a.Normalize()
	if n != e {
		t.Logf("Normalize(%+v) != %+v", a, e)
		t.Fail()
	}
}

func TestDot(t *testing.T) {
	e := 20.0
	x := Tuple{X: 1.0, Y: 2.0, Z: 3.0, W: 0.0}
	y := Tuple{X: 2.0, Y: 3.0, Z: 4.0, W: 0.0}
	d, _ := Dot(x, y)
	if d != e {
		t.Logf("Dot(%+v, %+v) != %f", x, y, e)
		t.Fail()
	}
}
