package raytracer

import (
	"math"
	"testing"
)

func TestNormalAt(t *testing.T) {
	ss := []Sphere{
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{0.0, 0.0, 0.0, 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{0.0, 0.0, 0.0, 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{0.0, 0.0, 0.0, 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{0.0, 0.0, 0.0, 1.0},
			Transform: IdentityMatrix(4),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{0.0, 0.0, 0.0, 1.0},
			Transform: TranslationMatrix(0, 1, 0),
		},
		Sphere{
			Radius:    1.0,
			Origin:    Tuple{0.0, 0.0, 0.0, 1.0},
			Transform: ScalingMatrix(1, .5, 1),
		},
	}

	sqrt2 := math.Sqrt(2) / 2
	sqrt3 := math.Sqrt(3) / 3
	ts := []Tuple{
		Tuple{1, 0, 0, 1},
		Tuple{0, 1, 0, 1},
		Tuple{0, 0, 1, 1},
		Tuple{sqrt3, sqrt3, sqrt3, 1},
		Tuple{0, 1.70711, -.70711, 1},
		Tuple{0, sqrt2, -sqrt2, 1},
	}

	e := []Tuple{
		Tuple{1, 0, 0, 0},
		Tuple{0, 1, 0, 0},
		Tuple{0, 0, 1, 0},
		Tuple{sqrt3, sqrt3, sqrt3, 0},
		Tuple{0, 0.70711, -0.70711, 0},
		Tuple{0, .97014, -0.24254, 0},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			norm, err := ss[i].NormalAt(ts[i])
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			if !TupleEqual(norm, x) {
				t.Logf("%+v != %+v", norm, x)
				t.Fail()
			}
		})
	}
}
