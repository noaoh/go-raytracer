package raytracer

import (
        "math"
        "testing"
)

func TestHit(t *testing.T) {
	s := Sphere{Radius: 1.0, Origin: Tuple{X: 0.0, Y: 0.0, Z: 0.0, W: 1.0}}

	is := [][]Intersection{
		{
			Intersection{
				T:   1.0,
				Obj: s,
			},
			Intersection{
				T:   2.0,
				Obj: s,
			},
		},
		{
			Intersection{
				T:   -1.0,
				Obj: s,
			},
			Intersection{
				T:   1.0,
				Obj: s,
			},
		},
		{
			Intersection{
				T:   -2.0,
				Obj: s,
			},
			Intersection{
				T:   -1.0,
				Obj: s,
			},
		},
		{
			Intersection{
				T:   5.0,
				Obj: s,
			},
			Intersection{
				T:   7.0,
				Obj: s,
			},
			Intersection{
				T:   -3.0,
				Obj: s,
			},
			Intersection{
				T:   2.0,
				Obj: s,
			},
		},
	}

	e := []Intersection{
		Intersection{
			T:   1.0,
			Obj: s,
		},
		Intersection{
			T:   1.0,
			Obj: s,
		},
		Intersection{
			T:   math.MaxFloat64,
			Obj: Sphere{},
		},
		Intersection{
			T:   2.0,
			Obj: s,
		},
	}

	for i, x := range e {
		if !IntersectionEqual(Hit(is[i]), x) {
			t.Fail()
			t.Logf("%+v != %+v", is[i], x)
		}
	}
}

func TestPrepareHit(t *testing.T) {
        rays := []Ray {
                Ray {
                        Origin : Tuple {0, 0, -5, 1},
                        Direction : Tuple {0, 0, 1, 0},
                },
        } 

        shapes := []Sphere {
                DefaultSphere(),
        }

        hits := []Intersection { 
                Intersection {
                        T: 4.0,
                        Obj: shapes[0],
                },
        }

        expected := []Intersection {
                Intersection {
                        T: 4.0,
                        Obj: shapes[0],
                        Point: Tuple{0, 0, -1, 1},
                        EyeV: Tuple{0, 0, -1, 0},
                        NormalV: Tuple{0, 0, -1, 0},
                        Inside: false,
                },
        }

        for i, x := range expected {
                t.Run(string(i), func(t *testing.T) {
                        err := hits[i].PrepareHit(rays[i]); if err != nil {
                                t.Log(err)
                                t.Fail()
                        }

                        if !IntersectionEqual(hits[i], x) {
                                t.Logf("%+v != %+v", hits[i], x)
                                t.Fail()
                        }
                })
        }
}

func TestShadeHit(t *testing.T) {
        worlds := []World {
                DefaultWorld(),
                DefaultWorld(),
        }

        worlds[1].Sources[0] = Light {
                Position: Tuple{0, .25, 0, 1},
                Intensity: Color{1, 1, 1},
        }

        shapes := []Sphere {
                worlds[0].Shapes[0],
                worlds[1].Shapes[1],
        }

        rays := []Ray {
                Ray {
                        Origin: Tuple{0, 0, -5, 1},
                        Direction: Tuple{0, 0, 1, 0},
                },
                Ray {
                        Origin: Tuple{0, 0, 0, 1},
                        Direction: Tuple{0, 0, 1, 0},
                },
        }
        
        hits := []Intersection {
                Intersection {
                        T: 4.0,
                        Obj: shapes[0],
                },
                Intersection {
                        T: 0.5,
                        Obj: shapes[1],
                },
        }

        expected := []Color {
                Color {.38066, .47583, .2855},
                Color {.90498, .90498, .90498},
        }

        for i, x := range expected {
                t.Run(string(i), func(t *testing.T) {
                        err := hits[i].PrepareHit(rays[i]); if err != nil {
                                t.Log(err)
                                t.Fail()
                        }

                        col, err := worlds[i].ShadeHit(hits[i]); if err != nil {
                                t.Log(err)
                                t.Fail()
                        }

                        if !ColorEqual(col, x) {
                                t.Logf("%+v != %+v", col, x)
                                t.Fail()
                        }
                })
        }
}

func TestColorAt(t *testing.T) {
        worlds := []World {
                DefaultWorld(),
                DefaultWorld(),
                DefaultWorld(),
        }

        worlds[2].Shapes[0].Material.Ambient = 1
        worlds[2].Shapes[1].Material.Ambient = 1

        rays := []Ray {
                Ray {
                        Origin: Tuple{X: 0, Y: 0, Z: -5, W:1},
                        Direction: Tuple{X: 0, Y: 1, Z: 0, W: 0},
                },
                Ray {
                        Origin: Tuple{X: 0, Y: 0, Z: -5, W:1},
                        Direction: Tuple{X: 0, Y: 0, Z: 1, W: 0},
                },
                Ray {
                        Origin: Tuple{X: 0, Y: 0, Z: -.75, W:1},
                        Direction: Tuple{X: 0, Y: 0, Z: 1, W: 0},
                },
        }

        expected := []Color {
                Color {R: 0.0, G: 0.0, B: 0.0},
                Color {R: 0.38066, G: 0.47583, B: 0.2855},
                Color {R: 1.0, G: 1.0, B: 1.0},
        }

        for i, x := range expected {
                t.Run(string(i), func(t *testing.T) {
                        col, err := worlds[i].ColorAt(rays[i]); if err != nil {
                                t.Log(err)
                                t.Fail()
                        }

                        if !ColorEqual(col, x) {
                                t.Logf("%+v != %+v", col, x)
                                t.Fail()
                        }
                })
        }
}
