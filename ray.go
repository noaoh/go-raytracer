package raytracer 

import (
	"fmt"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func CreateRay(origin, direction Tuple) (Ray, error) {
	if !origin.IsPoint() {
		return Ray{}, fmt.Errorf("origin must be point: %+v", origin)
	}

	if !direction.IsVector() {
		return Ray{}, fmt.Errorf("direction must be vector: %+v", direction)
	}

	return Ray{Origin: origin, Direction: direction}, nil
}

func (r *Ray) Position(t float64) (Tuple, error) {
	slope := r.Direction.MultiplyFloat(t)
	result, err := slope.Add(r.Origin)
	if err != nil {
		return Tuple{}, err
	}
	return result, nil
}

func Transform(r Ray, m Matrix) (Ray, error) {
        o, err := m.MultiplyTuple(r.Origin); if err != nil {
               return Ray{}, err 
        }

        d, err := m.MultiplyTuple(r.Direction); if err != nil {
               return Ray{}, err 
        }

        c, err := CreateRay(o, d); if err != nil {
               return Ray{}, err 
        }

        return c, nil
}

func RayEqual(r1, r2 Ray) bool {
        return TupleEqual(r1.Origin, r2.Origin) && TupleEqual(r1.Direction, r2.Direction)
}
