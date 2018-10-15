package raytracer 

import (
	"testing"
)

func TestPosition(t *testing.T) {
	ray, _ := CreateRay(Tuple{X: 2, Y: 3, Z: 4, W: 1}, Tuple{X: 1, Y: 0, Z: 0, W: 0})
	a := []float64{0.0, 1.0, -1.0, 2.5}
	e := []Tuple{
		Tuple{
			X: 2,
			Y: 3,
			Z: 4,
			W: 1,
		},
		Tuple{
			X: 3,
			Y: 3,
			Z: 4,
			W: 1,
		},
		Tuple{
			X: 1,
			Y: 3,
			Z: 4,
			W: 1,
		},
		Tuple{
			X: 4.5,
			Y: 3,
			Z: 4,
			W: 1,
		},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			p, err := ray.Position(a[i]); if err != nil {
                                t.Fail()
				t.Log(err)
			}

			if !TupleEqual(p, x) {
				t.Fail()
				t.Logf("%+v != %+v", p, x)
			}
		})
	}
}

func TestTransform(t *testing.T) {
        rs := []Ray {
                Ray{Tuple{1, 2, 3, 1}, Tuple{0, 1, 0, 0}},
                Ray{Tuple{1, 2, 3, 1}, Tuple{0, 1, 0, 0}},
        }
        
        ms := []Matrix {
                TranslationMatrix(3, 4, 5),
                ScalingMatrix(2, 3, 4),
        }

        e := []Ray {
                Ray{Tuple{4, 6, 8, 1}, Tuple{0, 1, 0, 0}},
                Ray{Tuple{2, 6, 12, 1}, Tuple{0, 3, 0, 0}},
        }

        for i, exp := range e {
                t.Run(string(i), func(t *testing.T) {
                        r, err := Transform(rs[i], ms[i]); if err != nil {
                                t.Fail()
                                t.Log(err)
                        }

                        if !RayEqual(r, exp) {
                                t.Fail()
                                t.Logf("%+v != %+v", r, exp)
                        }
                })
        }
}
