package raytracer

import (
	"math"
	"testing"
)

func TestIntersect(t *testing.T) {
	ss := []Sphere{
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
			Transform: ScalingMatrix(2, 2, 2),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0},
			Transform: TranslationMatrix(5, 0, 0),
		},
	}

	rs := []Ray{
		Ray{
			Origin:    Tuple{0.0, 0.0, -5.0, 1.0},
			Direction: Tuple{0.0, 0.0, 1.0, 0.0},
		},
		Ray{
			Origin:    Tuple{0.0, 1.0, -5.0, 1.0},
			Direction: Tuple{0.0, 0.0, 1.0, 0.0},
		},
		Ray{
			Origin:    Tuple{0.0, 2.0, -5.0, 1.0},
			Direction: Tuple{0.0, 0.0, 1.0, 0.0},
		},
		Ray{
			Origin:    Tuple{0.0, 0.0, 0.0, 1.0},
			Direction: Tuple{0.0, 0.0, 1.0, 0.0},
		},
		Ray{
			Origin:    Tuple{0.0, 0.0, 5.0, 1.0},
			Direction: Tuple{0.0, 0.0, 1.0, 0.0},
		},
		Ray{
			Origin:    Tuple{0.0, 0.0, -5.0, 1.0},
			Direction: Tuple{0.0, 0.0, 1.0, 0.0},
		},
		Ray{
			Origin:    Tuple{0.0, 0.0, -5.0, 1.0},
			Direction: Tuple{0.0, 0.0, 1.0, 0.0},
		},
	}

	e := [][]Intersection{
		{
			Intersection{
				T:   4.0,
				Obj: ss[0],
			},
			Intersection{
				T:   6.0,
				Obj: ss[0],
			},
		},
		{
			Intersection{
				T:   5.0,
				Obj: ss[1],
			},
			Intersection{
				T:   5.0,
				Obj: ss[1],
			},
		},
		{
			Intersection{
				T:   math.MaxFloat64,
				Obj: Sphere{},
			},
			Intersection{
				T:   math.MaxFloat64,
				Obj: Sphere{},
			},
		},
		{
			Intersection{
				T:   -1.0,
				Obj: ss[3],
			},
			Intersection{
				T:   1.0,
				Obj: ss[3],
			},
		},
		{
			Intersection{
				T:   -6.0,
				Obj: ss[4],
			},
			Intersection{
				T:   -4.0,
				Obj: ss[4],
			},
		},
		{
			Intersection{
				T:   3.0,
				Obj: ss[5],
			},
			Intersection{
				T:   7.0,
				Obj: ss[5],
			},
		},
		{
			Intersection{
				T:   math.MaxFloat64,
				Obj: Sphere{},
			},
			Intersection{
				T:   math.MaxFloat64,
				Obj: Sphere{},
			},
		},
	}

	for idx, iexp := range e {
		t.Run(string(idx), func(t *testing.T) {
			iact, err := ss[idx].Intersect(rs[idx])
			if err != nil {
				t.Fail()
				t.Log(err)
			}

			for iidx, x := range iexp {
				if !IntersectionEqual(iact[iidx], x) {
					t.Fail()
					t.Logf("%+v != %+v", iact[iidx], x)
				}
			}
		})
	}
}
