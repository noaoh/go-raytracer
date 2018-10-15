package raytracer 

import (
        "math"
)

type Intersection struct {
        T float64
        Obj Sphere
}

func IsPositive(f float64) bool {
        return math.Abs(f) == f
}

func Intersect(s Sphere, r Ray) ([]Intersection, error) {

        errValue := []Intersection{
                Intersection{T: math.MaxFloat64, Obj: Sphere {}},
                Intersection{T: math.MaxFloat64, Obj: Sphere {}},
        }
        

        invTrans, err := s.Transform.Inverse(); if err != nil {
                return errValue, err
        }

        r2, err := Transform(r, invTrans); if err != nil {
                return errValue, err
        }
        // r.Origin - Point(0, 0, 0)
        sphereToRay := Tuple {r2.Origin.X, r2.Origin.Y, r2.Origin.Z, 0}

        a, err := Dot(r2.Direction, r2.Direction); if err != nil {
                return errValue, err
        }

        x, err := Dot(r2.Direction, sphereToRay); if err != nil {
                return errValue, err
        }

        b := 2.0 * x

        // Already checked if sphereToRay is Point above
        y, _ := Dot(sphereToRay, sphereToRay)
        c := y - 1.0

        discriminant := (b * b) - (4 * a * c)

        // No real solutions exist
        if discriminant < 0 {
                return errValue, nil
        }

        t1 := (-1.0 * b - math.Sqrt(discriminant)) / (2.0 * a)
        t2 := (-1.0 * b + math.Sqrt(discriminant)) / (2.0 * a)
        if t1 > t2 {
                tmp := t1
                t1 = t2
                t2 = tmp
        }

        return []Intersection{{T: t1, Obj: s}, {T: t2, Obj: s}}, nil
}

func IntersectionEqual(i1, i2 Intersection) bool {
        return (i1.T == i2.T) && (SphereEqual(i1.Obj, i2.Obj))
}

func Hit(is []Intersection) Intersection {
        // Hit is always the lowest non-negative intersection
        low := Intersection{T: math.MaxFloat64, Obj: Sphere {}}
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
