package raytracer

import (
        "math"
)

func Hit(is []Intersection) Intersection {
	// Hit is always the lowest non-negative intersection
	low := Intersection{T: math.MaxFloat64, Obj: Sphere{}}
	for _, x := range is {
		// Ignore negative intersections
		if !IsPositive(x.T) {
			continue
		}

		if x.T < low.T {
                        low = x
		}
	}
	return low
}

func (hit *Intersection) PrepareHit(ray Ray) error {
        var err error
        hit.Point, err = ray.Position(hit.T); if err != nil {
                return err
        }

        hit.EyeV = ray.Direction.Negate()

        hit.NormalV, err = hit.Obj.NormalAt(hit.Point); if err != nil {
                return err
        }

        var normalEyeDot float64
        normalEyeDot, err = Dot(hit.NormalV, hit.EyeV); if err != nil {
                return err
        }

        if normalEyeDot < 0 {
                hit.Inside = true
                hit.NormalV = hit.NormalV.Negate()
        } else {
                hit.Inside = false
        }

        return nil
}

func (world World) ShadeHit(hit Intersection) (Color, error) {
        var c Color 
        for _, l := range world.Sources {
                col, err := Lighting(hit.Obj.Material, l, hit.Point, hit.EyeV, hit.NormalV)
                if err != nil {
                        return c, err
                }

                c = c.Add(col)
        }
        
        c = Color {R: Clamp(c.R, 0, 1), G: Clamp(c.G, 0, 1), B: Clamp(c.B, 0, 1)}
        return c, nil
}

func AllMisses(intersects []Intersection) (bool) {
        miss := Intersection{T: math.MaxFloat64, Obj: Sphere{}}
        for _, x := range intersects {
                if !IntersectionEqual(x, miss) {
                        return false
                }
        }

        return true

}
func (world World) ColorAt(ray Ray) (Color, error) {
        var c Color
        intersections, err := world.Intersect(ray); if err != nil {
                return c, err
        }
 
        if AllMisses(intersections) {
                return Color{0, 0, 0}, nil
        }

        hit := Hit(intersections)

        err = (&hit).PrepareHit(ray); if err != nil {
                return c, err
        }

        col, err := world.ShadeHit(hit); if err != nil {
                return c, err 
        }

        return col, nil

}
