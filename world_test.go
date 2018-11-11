package raytracer

import (
	"testing"
)

func TestWorldIntersect(t *testing.T) {
	world := DefaultWorld()
	ray := Ray{Tuple{0, 0, -5, 1}, Tuple{0, 0, 1, 0}}
	xs, err := world.Intersect(ray)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	s1 := DefaultSphere()
	s1.Material.Color = Color{R: 0.8, G: 1.0, B: 0.6}
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := DefaultSphere()
	s2.Transform = ScalingMatrix(0.5, 0.5, 0.5)
	e := []Intersection{
		Intersection{
			T:   4,
			Obj: s1,
		},
		Intersection{
			T:   4.5,
			Obj: s2,
		},
		Intersection{
			T:   5.5,
			Obj: s2,
		},
		Intersection{
			T:   6,
			Obj: s1,
		},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			if !IntersectionEqual(xs[i], x) {
				t.Logf("%+v != %+v", xs[i], x)
				t.Fail()
			}
		})
	}
}
