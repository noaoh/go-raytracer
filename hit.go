package raytracer

import (
_       "log"
	"math"
)

func Hit(is []Intersection) (Intersection, bool) {
	// Hit is always the lowest non-negative intersection
	low := Intersection{T: math.MaxFloat64, Obj: Sphere{}}
	for _, x := range is {
		// Ignore negative intersections
		if IsPositive(x.T) && x.T < low.T {
			low = x
		}
	}

        if low.T < math.MaxFloat64 {
                return low, true
        } else {
                return low, false
        }
}

func (hit *Intersection) PrepareHit(ray Ray) error { 
        var err error
	hit.Point, err = ray.Position(hit.T)
	if err != nil {
		return err
	}

	hit.EyeV = ray.Direction.Negate()

	hit.NormalV, err = hit.Obj.NormalAt(hit.Point)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	var normalEyeDot float64
	normalEyeDot, err = Dot(hit.NormalV, hit.EyeV)
	if err != nil {
		return err
	}

        hit.Inside = normalEyeDot < 0
	if hit.Inside {
		hit.NormalV = hit.NormalV.Negate()
	}

        // move point slightly above surface, into direction of normal, to make
        // shadows render properly
	hit.Point, err = hit.Point.Add(hit.NormalV.MultiplyFloat(.00001))
	return nil
}

func (world World) ShadeHit(hit Intersection) (Color, error) {
        shadowed, err := world.isShadowed(hit.Point)
        if err != nil {
                return Color{}, err
        }

        col, err := Lighting(hit.Obj.Material, world.Sources[0], hit.Point, hit.EyeV, hit.NormalV, shadowed)
        if err != nil {
                return col, err
        }

	return col, nil
}

func AllMisses(intersects []Intersection) bool {
	miss := Intersection{T: math.MaxFloat64, Obj: Sphere{}}
	for _, x := range intersects {
		if !IntersectionEqual(x, miss) {
			return false
		}
	}

	return true

}
func (world World) ColorAt(ray Ray) (Color, error) {
        black := Color{0, 0, 0}
	intersections, err := world.Intersect(ray)
	if err != nil {
		return black, err
	}

	hit, hasHit := Hit(intersections)
        if hasHit {
                err := (&hit).PrepareHit(ray)
                if err != nil {
                        return black, err
                }
                
                res, err := world.ShadeHit(hit)
                if err != nil {
                        return black, err
                }

                return res, nil
        }

	return black, nil
}
