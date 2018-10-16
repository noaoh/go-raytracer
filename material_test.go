package raytracer

import (
	"math"
	"testing"
)

func TestLighting(t *testing.T) {
	pos := Tuple{0.0, 0.0, 0.0, 1.0}

	m := DefaultMaterial()

	sqrt2 := math.Sqrt(2) / 2
	es := []Tuple{
		Tuple{0.0, 0.0, -1.0, 0.0},
		Tuple{0.0, sqrt2, -sqrt2, 0.0},
		Tuple{0.0, 0.0, -1.0, 0.0},
		Tuple{0.0, -sqrt2, -sqrt2, 0.0},
		Tuple{0.0, 0.0, -1.0, 0.0},
	}

	ns := []Tuple{
		Tuple{0.0, 0.0, -1.0, 0.0},
		Tuple{0.0, 0.0, -1.0, 0.0},
		Tuple{0.0, 0.0, -1.0, 0.0},
		Tuple{0.0, 0.0, -1.0, 0.0},
		Tuple{0.0, 0.0, -1.0, 0.0},
	}

	ls := []Light{
		Light{
			Position:  Tuple{0.0, 0.0, -10.0, 1.0},
			Intensity: Color{1.0, 1.0, 1.0},
		},
		Light{
			Position:  Tuple{0.0, 0.0, -10.0, 1.0},
			Intensity: Color{1.0, 1.0, 1.0},
		},
		Light{
			Position:  Tuple{0.0, 10.0, -10.0, 1.0},
			Intensity: Color{1.0, 1.0, 1.0},
		},
		Light{
			Position:  Tuple{0.0, 10.0, -10.0, 1.0},
			Intensity: Color{1.0, 1.0, 1.0},
		},
		Light{
			Position:  Tuple{0.0, 0.0, 10.0, 1.0},
			Intensity: Color{1.0, 1.0, 1.0},
		},
	}

	e := []Color{
		Color{1.9, 1.9, 1.9},
		Color{1.0, 1.0, 1.0},
		Color{.7364, .7364, .7364},
		Color{1.6364, 1.6364, 1.6364},
		Color{0.1, 0.1, 0.1},
	}

	for i, x := range e {
		t.Run(string(i), func(t *testing.T) {
			c, err := Lighting(m, ls[i], pos, es[i], ns[i])
			if err != nil {
				t.Log(err)
				t.Fail()
			}

			if !ColorEqual(c, x) {
				t.Logf("%+v != %+v", c, x)
				t.Fail()
			}
		})
	}
}
