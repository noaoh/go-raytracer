package matrix

import (
	tup "github.com/noaoh/raytracer/models/tuple"
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
			z := identityMatrix(a[i])
			if !Equal(z, x) {
				t.Logf("%+v != %+v", z, x)
				t.Fail()
			}
		})
	}
}

func TestMultiply(t *testing.T) {
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
			},
			Matrix{
				Rows: 4,
				Cols: 4,
				Data: [][]float64{
					{0.0, 1.0, 2.0, 4},
					{1.0, 2.0, 4.0, 8},
					{2.0, 4.0, 8.0, 16},
					{4.0, 8.0, 16.0, 32},
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
			fromTuple(tup.Tuple{X: 1.0, Y: 2.0, Z: 3.0, W: 1.0}),
		},
	}

	e := []Matrix{
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{24, 49, 98, 196},
				{31, 64, 128, 256},
				{38, 79, 158, 316},
				{45, 94, 188, 376},
			},
		},
		Matrix{
			Rows: 4,
			Cols: 1,
			Data: [][]float64{{18}, {24}, {33}, {1}},
		},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			j := a[i][0]
			k := a[i][1]
			y, _ := Multiply(j, k); if !Equal(y, x) {
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
                                {1, 5},
                                {-3, 2},
                        },
                },
		Matrix{
			Rows: 2,
			Cols: 2,
			Data: [][]float64{
				{5, 0},
				{-1, 5},
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
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{-4, 2, -2, 3},
				{9, 6, 2, 6},
				{0, -5, 1, -5},
				{0, 0, 0, 0},
			},
		},
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
	}

	e := []float64{17, 25, -196, -4071, 0, 532}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			d, _ := a[i].Determinant()
			if d != x {
				t.Logf("%+v != %+v", d, x)
				t.Fail()
			}
		})
	}
}

func TestSubmatrix(t *testing.T) {
	in := [][]int{{1, 0}, {0, 2}, {2, 1}}

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
				{1, 5, 0},
				{-3, 2, 7},
				{0, 6, -3},
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
			if !Equal(s, x) {
				t.Logf("%+v != %+v", s, x)
				t.Fail()
			}
		})
	}
}

func TestMinor(t *testing.T) {
        a := []Matrix {
                Matrix {
                        Rows: 3,
                        Cols: 3,
                        Data: [][]float64 {
                                {3, 5, 0},
                                {2, -1, -7},
                                {6, -1, 5},
                        },
                },
                Matrix {
                        Rows: 3,
                        Cols: 3,
                        Data: [][]float64 {
                                {1, 2, 6},
                                {-5, 8, -4},
                                {2, 6, 4},
                        },
                },
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                                {-2, -8, 3, 5},
                                {-3, 1, 7, 3},
                                {1, 2, -9, 6},
                                {-6, 7, 7, -9},
                        },
                },
        }

        e := []float64 {25, -46, -51}

        in := [][]int {{1, 0}, {0, 2}, {0, 3}}

        for i, x := range e {
                t.Run(string(i), func(t *testing.T) {
                        c := a[i].Minor(in[i][0], in[i][1]); if c != x {
                                t.Logf("%+v != %+v", c, x)
                                t.Fail()
                        }
                })
        }
}

func TestCofactor(t *testing.T) {
        a := []Matrix {
                Matrix {
                        Rows: 3,
                        Cols: 3,
                        Data: [][]float64 {
                                {3, 5, 0},
                                {2, -1, -7},
                                {6, -1, 5},
                        },
                },
                Matrix {
                        Rows: 3,
                        Cols: 3,
                        Data: [][]float64 {
                                {1, 2, 6},
                                {-5, 8, -4},
                                {2, 6, 4},
                        },
                },
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                                {-2, -8, 3, 5},
                                {-3, 1, 7, 3},
                                {1, 2, -9, 6},
                                {-6, 7, 7, -9},
                        },
                },
        }

        e := []float64 {-25, -46, 51}

        in := [][]int {{1, 0}, {0, 2}, {0, 3}}

        for i, x := range e {
                t.Run(string(i), func(t *testing.T) {
                        c := a[i].Cofactor(in[i][0], in[i][1]); if c != x {
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
		identityMatrix(4),
	}

	e := []Matrix{
		Matrix{
			Rows: 4,
			Cols: 4,
			Data: [][]float64{
				{0, 9, 1, 0},
				{9, 8, 8, 0},
				{3, 0, 5, 5},
				{0, 8, 3, 3},
			},
		},
		identityMatrix(4),
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			tr := Transpose(a[i])
			if !Equal(tr, x) {
				t.Logf("%+v != %+v", tr, x)
			}
		})
	}
}

func TestInverse(t *testing.T) {
        a := []Matrix {
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                                {6, 4, 4, 4},
                                {5, 5, 6, 6},
                                {4, -9, 3, -7},
                                {9, 1, 7, -6},
                        },
                },
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                                {-5, 2, 6, -8},
                                {1, -5, 1, 8},
                                {7, 7, -6, -7},
                                {1, -3, 7, 4},
                        },
                },
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                                {8, -5, 9, 2},
                                {7, 5, 6, 1},
                                {-6, 0, 9, 6},
                                {-3, 0, -9, -4},
                        },
                },
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                            { 9, 3, 0, 9},
                            {-5, -2, -6, -3},
                            {-4, 9, 6, 4},
                            {-7, 6, 6, 2},
                        },
                },
        }

        e := []Matrix {
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                                {607.0/1570, -192.0/785, 26.0/785, -4.0/157},
                                {-73.0/1570, -17.0/785, -104.0/785, 16.0/157},
                                {-62.0/157, 55.0/157, 4.0/157, 9.0/157},
                                {35.0/314, 6.0/157, 9.0/157, -19.0/157},
                        },
                        

                },
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                                {29.0/133, 60.0/133, 32.0/133, -6.0/133},
                                {-215.0/266, -775.0/532, -59.0/133, 277.0/532},
                                {-3.0/38, -17.0/76, -1.0/19, 15.0/76},
                                {-139.0/266, -433.0/532, -40.0/133, 163.0/532},
                        },
                },
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                            {-0.15385, -0.15385, -0.28205, -0.53846},
                            {-0.07692, 0.12308, 0.02564, 0.03077},
                            { 0.35897, 0.35897, 0.43590, 0.92308},
                            {-0.69231, -0.69231, -0.76923, -1.92308},
                        },
                },
                Matrix {
                        Rows: 4,
                        Cols: 4,
                        Data: [][]float64 {
                            {-0.04074, -0.07778, 0.14444, -0.22222},
                            {-0.07778, 0.03333, 0.36667, -0.33333},
                            {-0.02901, -0.14630, -0.10926, 0.12963},
                            { 0.17778, 0.06667, -0.26667, 0.33333},
                    },
                },
        }

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			inv, _ := a[i].Inverse(); if !Equal(inv, x) {
				t.Logf("%+v != %+v", inv, x)
			}
		})
	}
}
