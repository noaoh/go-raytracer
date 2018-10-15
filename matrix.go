package raytracer 

import (
	"fmt"
	"math"
)

type Matrix struct {
	Rows int
	Cols int
	Data [][]float64
}

func FloatEqual(a, b float64) bool {
	epsilon := .00001
	return (math.Abs(a-b) < epsilon)
}

func CreateMatrix(r, c int) Matrix {
	m := Matrix{Rows: r, Cols: c}
	m.Data = make([][]float64, r)
	for x := range m.Data {
		m.Data[x] = make([]float64, c)
	}
	return m
}

func IdentityMatrix(r int) Matrix {
	m := CreateMatrix(r, r)
	for i := 0; i < r; i++ {
		m.Data[i][i] = 1
	}
	return m
}

func (m Matrix) Row(x int) []float64 {
	return m.Data[x]
}

func (m Matrix) Col(y int) []float64 { cols := make([]float64, y)
	for i := 0; i < m.Cols; i++ {
		cols[i] = m.Data[i][y]
	}

	return cols
}

func (a Matrix) MultiplyMatrix(b Matrix) (Matrix, error) {
	if a.Cols != b.Rows {
		return Matrix{}, fmt.Errorf("The number of columns in matrix a must be the same as the number of rows in matrix b for matrix multiplication: a = %+v, b = %+v", a, b)
	}

	m := CreateMatrix(a.Rows, b.Cols)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			var sum float64 = 0
			for k := 0; k < a.Cols; k++ {
				sum += a.Data[i][k] * b.Data[k][j]
			}
			m.Data[i][j] = sum
		}
	}

	return m, nil
}

func (a *Matrix) MultiplyFloat(f float64) Matrix {
	m := CreateMatrix(a.Rows, a.Cols)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < a.Cols; j++ {
			m.Data[i][j] = a.Data[i][j] * f
		}
	}

	return m
}

func (a *Matrix) MultiplyTuple(b Tuple) (Tuple, error) {
	m := FromTuple(b)

	r, err := a.MultiplyMatrix(m)
	if err != nil {
		return Tuple{}, err
	}

	t, err := FromMatrix(r)
	if err != nil {
		return Tuple{}, err
	}

	return t, nil
}

func FromTuple(tup Tuple) Matrix {
	m := CreateMatrix(4, 1)
	m.Data[0][0] = tup.X
	m.Data[1][0] = tup.Y
	m.Data[2][0] = tup.Z
	m.Data[3][0] = tup.W
	return m
}

func Transpose(m Matrix) Matrix {
	r := CreateMatrix(m.Cols, m.Rows)

	for y, s := range m.Data {
		for x, e := range s {
			r.Data[x][y] = e
		}
	}

	return r
}

func (m Matrix) SubMatrix(row, col int) Matrix {
	res := CreateMatrix(m.Rows-1, m.Cols-1)

	ri := 0
	for i, r := range m.Data {
		if i == row {
			continue
		}

		rj := 0
		for j, d := range r {
			if j == col {
				continue
			}

			res.Data[ri][rj] = d
			rj += 1
		}
		ri += 1
	}
	return res
}

func (m Matrix) Minor(row, col int) float64 {
	// The SubMatrix of a square matrix is also square
	d, _ := m.SubMatrix(row, col).Determinant()
	return d
}

func (m Matrix) Cofactor(row, col int) float64 {
	return math.Pow(-1.0, float64(col+row)) * m.Minor(row, col)
}

func (m Matrix) Determinant() (float64, error) {
	res := 0.0
	if m.Rows != m.Cols {
		return 1, fmt.Errorf("Matrix must have same number of columns and rows to compute determinant: %+v", m)
	} else if m.Rows == 2 {
		return m.Data[0][0]*m.Data[1][1] - m.Data[0][1]*m.Data[1][0], nil
	} else {
		for j := 0; j < m.Cols; j++ {
			res += m.Cofactor(0, j) * m.Data[0][j]
		}
	}
	return res, nil
}

func (m Matrix) Inverse() (Matrix, error) {
	if m.Rows != m.Cols {
		return Matrix{}, fmt.Errorf("Matrix must have same number of columns and rows to compute inverse: %+v", m)
	}

	d, _ := m.Determinant()

	if d == 0 {
		return Matrix{}, fmt.Errorf("Determinant of matrix must not be zero: %+v", m)
	}

	r := CreateMatrix(m.Rows, m.Cols)
	for i := 0; i < r.Rows; i++ {
		for j := 0; j < r.Cols; j++ {
			r.Data[i][j] = m.Cofactor(i, j) / float64(d)
		}
	}

	return Transpose(r), nil
}

func MatrixEqual(m, n Matrix) bool {
	if m.Rows != n.Rows || m.Cols != n.Cols {
		return false
	}

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if !FloatEqual(m.Data[i][j], n.Data[i][j]) {
				return false
			}
		}
	}

	return true
}

func TranslationMatrix(x, y, z float64) Matrix {
	m := IdentityMatrix(4)
	m.Data[0][3] = x
	m.Data[1][3] = y
	m.Data[2][3] = z
	return m
}

func ScalingMatrix(x, y, z float64) Matrix {
	m := IdentityMatrix(4)
	m.Data[0][0] = x
	m.Data[1][1] = y
	m.Data[2][2] = z
	return m
}

func XAxisRotationMatrix(r float64) Matrix {
	m := IdentityMatrix(4)
	cos := math.Cos(r)
	sin := math.Sin(r)

	m.Data[1][1] = cos
	m.Data[1][2] = -1 * sin
	m.Data[2][1] = sin
	m.Data[2][2] = cos
	return m
}

func YAxisRotationMatrix(r float64) Matrix {
	m := IdentityMatrix(4)
	cos := math.Cos(r)
	sin := math.Sin(r)

	m.Data[0][0] = cos
	m.Data[0][2] = sin
	m.Data[2][0] = -1 * sin
	m.Data[2][2] = cos
	return m
}

func ZAxisRotationMatrix(r float64) Matrix {
	m := IdentityMatrix(4)
	cos := math.Cos(r)
	sin := math.Sin(r)

	m.Data[0][0] = cos
	m.Data[0][1] = -1 * sin
	m.Data[1][0] = sin
	m.Data[1][1] = cos
	return m
}

func ShearingMatrix(xy, xz, yx, yz, zx, zy float64) Matrix {
	m := IdentityMatrix(4)

	m.Data[0][1] = xy
	m.Data[0][2] = xz
	m.Data[1][0] = yx
	m.Data[1][2] = yz
	m.Data[2][0] = zx
	m.Data[2][1] = zy
	return m
}
