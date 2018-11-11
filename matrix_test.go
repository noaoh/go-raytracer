package raytracer

import (
	"math"
	"testing"
)

func TestIdentityMatrix(t *testing.T) {
	a := []int{1, 2, 3, 4}
	e := []Matrix{
		Matrix{
			Rows: 1,
			Cols: 1,
			Data: [][]float64{{1.0}},
		},
		Matrix{
			Rows: 2,
			Cols: 2,
			Data: [][]float64{
				{1.0, 0.0},
				{0.0, 1.0},
			},
		},
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{1.0, 0.0, 0.0},
				{0.0, 1.0, 0.0},
				{0.0, 0.0, 1.0},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{1.0, 0.0, 0.0, 0.0},
				{0.0, 1.0, 0.0, 0.0},
				{0.0, 0.0, 1.0, 0.0},
				{0.0, 0.0, 0.0, 1.0},
			},
		},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			z := IdentityMatrix(a[i])
			if !MatrixEqual(z, x) {
				t.Logf("%+v != %+v", z, x)
				t.Fail()
			}
		})
	}
}

func TestMultiplyMatrix(t *testing.T) {
	a := [][]Matrix{
		{
			Matrix{
				Rows: 4,
				Cols: 4,
				Data: [][]float64{
					{1.0, 2.0, 3.0, 4.0},
					{2.0, 3.0, 4.0, 5.0},
					{3.0, 4.0, 5.0, 6.0},
					{4.0, 5.0, 6.0, 7.0},
				},
			}, Matrix{
				Rows: 4,
				Cols: 4,
				Data: [][]float64{
					{0.0, 1.0, 2.0, 4.0},
					{1.0, 2.0, 4.0, 8.0},
					{2.0, 4.0, 8.0, 16.0},
					{4.0, 8.0, 16.0, 32.0},
				},
			},
		},
		{
			Matrix{
				Rows: 4,
				Cols: 4,
				Data: [][]float64{
					{1.0, 2.0, 3.0, 4.0},
					{2.0, 4.0, 4.0, 2.0},
					{8.0, 6.0, 4.0, 1.0},
					{0.0, 0.0, 0.0, 1.0},
				},
			},
			FromTuple(Tuple{X: 1.0, Y: 2.0, Z: 3.0, W: 1.0}),
		},
	}

	e := []Matrix{
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{24.0, 49.0, 98.0, 196.0},
				{31.0, 64.0, 128.0, 256.0},
				{38.0, 79.0, 158.0, 316.0},
				{45.0, 94.0, 188.0, 376.0},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 1,
			Data: [][]float64{{18.0}, {24.0}, {33.0}, {1.0}},
		},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			j := a[i][0]
			k := a[i][1]
			y, _ := j.MultiplyMatrix(k)
			if !MatrixEqual(y, x) {
				t.Logf("%+v != %+v\n", y, x)
				t.Fail()
			}
		})
	}

}

func TestDeterminant(t *testing.T) {
	a := []Matrix{
		Matrix{
			Rows: 2,
			Cols: 2,
			Data: [][]float64{
				{1.0, 5.0},
				{-3.0, 2.0},
			},
		},
		Matrix{
			Rows: 2,
			Cols: 2,
			Data: [][]float64{
				{5.0, 0.0},
				{-1.0, 5.0},
			},
		},
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{1.0, 2.0, 6.0},
				{-5.0, 8.0, -4.0},
				{2.0, 6.0, 4.0},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-2.0, -8.0, 3.0, 5.0},
				{-3.0, 1.0, 7.0, 3.0},
				{1.0, 2.0, -9.0, 6.0},
				{-6.0, 7.0, 7.0, -9.0},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-4.0, 2.0, -2.0, 3},
				{9.0, 6.0, 2.0, 6},
				{0.0, -5.0, 1.0, -5},
				{0.0, 0.0, 0.0, 0},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-5.0, 2.0, 6.0, -8},
				{1.0, -5.0, 1.0, 8},
				{7.0, 7.0, -6.0, -7},
				{1.0, -3.0, 7.0, 4},
			},
		},
	}

	e := []float64{17.0, 25.0, -196.0, -4071.0, 0.0, 532}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			d, _ := a[i].Determinant()
			if !FloatEqual(d, x) {
				t.Logf("%+v != %+v", d, x)
				t.Fail()
			}
		})
	}
}

func TestSubmatrix(t *testing.T) {
	in := [][]int{{1.0, 0}, {0.0, 2}, {2.0, 1}}

	a := []Matrix{
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{3.0, 5.0, 0},
				{2.0, -1.0, -7},
				{6.0, -1.0, 5},
			},
		},
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{1.0, 5.0, 0},
				{-3.0, 2.0, 7},
				{0.0, 6.0, -3},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-6, 1, 1, 6},
				{-8, 5, 8, 6},
				{-1, 0, 8, 2},
				{-7, 1, -1, 1},
			},
		},
	}

	e := []Matrix{
		Matrix{
			Rows: 2,
			Cols: 2,
			Data: [][]float64{
				{5, 0},
				{-1, 5},
			},
		},
		Matrix{
			Rows: 2,
			Cols: 2,
			Data: [][]float64{
				{-3, 2},
				{0, 6},
			},
		},
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{-6, 1, 6},
				{-8, 8, 6},
				{-7, -1, 1},
			},
		},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			s := a[i].SubMatrix(in[i][0], in[i][1])
			if !MatrixEqual(s, x) {
				t.Logf("%+v != %+v", s, x)
				t.Fail()
			}
		})
	}
}

func TestMinor(t *testing.T) {
	a := []Matrix{
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
		},
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-2, -8, 3, 5},
				{-3, 1, 7, 3},
				{1, 2, -9, 6},
				{-6, 7, 7, -9},
			},
		},
	}

	e := []float64{25, -46, -51}

	in := [][]int{{1, 0}, {0, 2}, {0, 3}}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			c := a[i].Minor(in[i][0], in[i][1])
			if !FloatEqual(c, x) {
				t.Logf("%+v != %+v", c, x)
				t.Fail()
			}
		})
	}
}

func TestCofactor(t *testing.T) {
	a := []Matrix{
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
		},
		Matrix{
			Rows: 3,
			Cols: 3,
			Data: [][]float64{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-2, -8, 3, 5},
				{-3, 1, 7, 3},
				{1, 2, -9, 6},
				{-6, 7, 7, -9},
			},
		},
	}

	e := []float64{-25, -46, 51}

	in := [][]int{{1, 0}, {0, 2}, {0, 3}}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			c := a[i].Cofactor(in[i][0], in[i][1])
			if !FloatEqual(c, x) {
				t.Logf("%+v != %+v", c, x)
				t.Fail()
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	a := []Matrix{
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{0, 9, 3, 0},
				{9, 8, 0, 8},
				{1, 8, 5, 3},
				{0, 0, 5, 8},
			},
		},
		IdentityMatrix(4),
	}

	e := []Matrix{
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{0, 9, 1, 0},
				{9, 8, 8, 0},
				{3, 0, 5, 5},
				{0, 8, 3, 8},
			},
		},
		IdentityMatrix(4),
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			tr := a[i].Transpose()
			if !MatrixEqual(tr, x) {
				t.Logf("%+v != %+v", tr, x)
				t.Fail()
			}
		})
	}
}

func TestInverse(t *testing.T) {
	a := []Matrix{
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-5, 2, 6, -8},
				{1, -5, 1, 8},
				{7, 7, -6, -7},
				{1, -3, 7, 4},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{8, -5, 9, 2},
				{7, 5, 6, 1},
				{-6, 0, 9, 6},
				{-3, 0, -9, -4},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{9, 3, 0, 9},
				{-5, -2, -6, -3},
				{-4, 9, 6, 4},
				{-7, 6, 6, 2},
			},
		},
	}

	e := []Matrix{
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{0.21805, 0.45113, 0.24060, -0.04511},
				{-0.80827, -1.45677, -0.44361, 0.52068},
				{-0.07895, -0.22368, -0.05263, 0.19737},
				{-0.52256, -0.81391, -0.30075, 0.30639},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-0.15385, -0.15385, -0.28205, -0.53846},
				{-0.07692, 0.12308, 0.02564, 0.03077},
				{0.35897, 0.35897, 0.43590, 0.92308},
				{-0.69231, -0.69231, -0.76923, -1.92308},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-0.04074, -0.07778, 0.14444, -0.22222},
				{-0.07778, 0.03333, 0.36667, -0.33333},
				{-0.02901, -0.14630, -0.10926, 0.12963},
				{0.17778, 0.06667, -0.26667, 0.33333},
			},
		},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			inv, _ := a[i].Inverse()
			if !MatrixEqual(inv, x) {
				t.Logf("%+v != %+v", inv, x)
				t.Fail()
			}
		})
	}
}

func TestTranslationMatrix(t *testing.T) {
	tr := TranslationMatrix(5, -3, 2)
	inv, _ := tr.Inverse()
	in := []Matrix{
		tr,
		inv,
		tr,
	}

	a := []Tuple{
		Tuple{X: -3, Y: 4, Z: 5, W: 1},
		Tuple{X: -3, Y: 4, Z: 5, W: 1},
		Tuple{X: -3, Y: 4, Z: 5, W: 0},
	}

	e := []Tuple{
		Tuple{X: 2, Y: 1, Z: 7, W: 1},
		Tuple{X: -8, Y: 7, Z: 3, W: 1},
		Tuple{X: -3, Y: 4, Z: 5, W: 0},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			v, _ := in[i].MultiplyTuple(a[i])
			if !TupleEqual(v, x) {
				t.Logf("%+v != %+v", v, x)
				t.Fail()
			}
		})
	}
}

func TestScalingMatrix(t *testing.T) {
	s := ScalingMatrix(2, 3, 4)
	inv, _ := s.Inverse()

	in := []Matrix{
		s,
		s,
		inv,
	}

	a := []Tuple{
		Tuple{X: -4, Y: 6, Z: 8, W: 1},
		Tuple{X: -4, Y: 6, Z: 8, W: 0},
		Tuple{X: -4, Y: 6, Z: 8, W: 0},
	}

	e := []Tuple{
		Tuple{X: -8, Y: 18, Z: 32, W: 1},
		Tuple{X: -8, Y: 18, Z: 32, W: 0},
		Tuple{X: -2, Y: 2, Z: 2, W: 0},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			v, _ := in[i].MultiplyTuple(a[i])
			if !TupleEqual(v, x) {
				t.Logf("%+v != %+v", v, x)
				t.Fail()
			}
		})
	}
}

func TestXAxisRotationMatrix(t *testing.T) {
	hq := XAxisRotationMatrix(math.Pi / 4)
	inv, _ := hq.Inverse()

	in := []Matrix{
		hq,
		XAxisRotationMatrix(math.Pi / 2),
		inv,
	}

	a := Tuple{X: 0, Y: 1, Z: 0, W: 1}

	e := []Tuple{
		Tuple{X: 0.0, Y: math.Sqrt(2) / 2.0, Z: math.Sqrt(2) / 2.0, W: 1.0},
		Tuple{X: 0.0, Y: 0.0, Z: 1.0, W: 1.0},
		Tuple{X: 0.0, Y: math.Sqrt(2) / 2.0, Z: math.Sqrt(2) / -2.0, W: 1.0},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			v, _ := in[i].MultiplyTuple(a)
			if !TupleEqual(v, x) {
				t.Logf("%+v != %+v", v, x)
				t.Fail()
			}
		})
	}
}

func TestYAxisRotationMatrix(t *testing.T) {
	hq := YAxisRotationMatrix(math.Pi / 4)
	inv, _ := hq.Inverse()

	in := []Matrix{
		hq,
		YAxisRotationMatrix(math.Pi / 2),
		inv,
	}

	a := Tuple{X: 0.0, Y: 0.0, Z: 1.0, W: 1.0}

	e := []Tuple{
		Tuple{X: math.Sqrt(2) / 2.0, Y: 0.0, Z: math.Sqrt(2) / 2.0, W: 1.0},
		Tuple{X: 1.0, Y: 0.0, Z: 0.0, W: 1.0},
		Tuple{X: math.Sqrt(2) / -2.0, Y: 0.0, Z: math.Sqrt(2) / 2.0, W: 1.0},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			v, _ := in[i].MultiplyTuple(a)
			if !TupleEqual(v, x) {
				t.Logf("%+v != %+v", v, x)
				t.Fail()
			}
		})
	}
}

func TestZAxisRotationMatrix(t *testing.T) {
	hq := ZAxisRotationMatrix(math.Pi / 4.0)
	inv, _ := hq.Inverse()

	in := []Matrix{
		hq,
		ZAxisRotationMatrix(math.Pi / 2.0),
		inv,
	}

	a := Tuple{X: 0.0, Y: 1.0, Z: 0.0, W: 1}

	e := []Tuple{
		Tuple{X: math.Sqrt(2) / -2.0, Y: math.Sqrt(2) / 2.0, Z: 0.0, W: 1.0},
		Tuple{X: -1.0, Y: 0.0, Z: 0.0, W: 1.0},
		Tuple{X: math.Sqrt(2) / 2.0, Y: math.Sqrt(2) / 2.0, Z: 0.0, W: 1.0},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			v, _ := in[i].MultiplyTuple(a)
			if !TupleEqual(v, x) {
				t.Logf("%+v != %+v", v, x)
				t.Fail()
			}
		})
	}
}

func TestShearingMatrix(t *testing.T) {
	in := []Matrix{
		ShearingMatrix(0.0, 1.0, 0.0, 0.0, 0.0, 0.0),
		ShearingMatrix(0.0, 0.0, 1.0, 0.0, 0.0, 0.0),
		ShearingMatrix(0.0, 0.0, 0.0, 1.0, 0.0, 0.0),
		ShearingMatrix(0.0, 0.0, 0.0, 0.0, 1.0, 0.0),
		ShearingMatrix(0.0, 0.0, 0.0, 0.0, 0.0, 1.0),
	}

	a := Tuple{X: 2.0, Y: 3.0, Z: 4.0, W: 1.0}

	e := []Tuple{
		Tuple{X: 6.0, Y: 3.0, Z: 4.0, W: 1.0},
		Tuple{X: 2.0, Y: 5.0, Z: 4.0, W: 1.0},
		Tuple{X: 2.0, Y: 7.0, Z: 4.0, W: 1.0},
		Tuple{X: 2.0, Y: 3.0, Z: 6.0, W: 1.0},
		Tuple{X: 2.0, Y: 3.0, Z: 7.0, W: 1.0},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			v, _ := in[i].MultiplyTuple(a)
			if !TupleEqual(v, x) {
				t.Logf("%+v != %+v", v, x)
				t.Fail()
			}
		})
	}
}

func TestViewTransform(t *testing.T) {
	froms := []Tuple{
		Tuple{X: 0, Y: 0, Z: 0, W: 1},
		Tuple{X: 0, Y: 0, Z: 0, W: 1},
		Tuple{X: 0, Y: 0, Z: 8, W: 1},
		Tuple{X: 1, Y: 3, Z: 2, W: 1},
	}

	tos := []Tuple{
		Tuple{X: 0, Y: 0, Z: -1, W: 1},
		Tuple{X: 0, Y: 0, Z: 1, W: 1},
		Tuple{X: 0, Y: 0, Z: 0, W: 1},
	}

	ups := []Tuple{
		Tuple{X: 0, Y: 1, Z: 0, W: 0},
		Tuple{X: 0, Y: 1, Z: 0, W: 0},
		Tuple{X: 0, Y: 1, Z: 0, W: 0},
	}

	expected := []Matrix{
		IdentityMatrix(4),
		ScalingMatrix(-1, 1, -1),
		TranslationMatrix(0, 0, -8),
	}

	for i, x := range expected {
		t.Run(string(i), func(t *testing.T) {
			v, err := ViewTransform(froms[i], tos[i], ups[i])
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			if !MatrixEqual(v, x) {
				t.Logf("%+v != %+v", v, x)
				t.Fail()
			}
		})
	}
}
